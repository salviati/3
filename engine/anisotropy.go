package engine

// Magnetocrystalline anisotropy.

import (
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
)

// Anisotropy variables
var (
	Ku1        = NewScalarParam("Ku1", "J/m3", "1st order uniaxial anisotropy constant", &ku1_red)
	Ku2        = NewScalarParam("Ku2", "J/m3", "2nd order uniaxial anisotropy constant", &ku2_red)
	Kc1        = NewScalarParam("Kc1", "J/m3", "1st order cubic anisotropy constant", &kc1_red)
	Kc2        = NewScalarParam("Kc2", "J/m3", "2nd order cubic anisotropy constant", &kc2_red)
	Kc3        = NewScalarParam("Kc3", "J/m3", "3rd order cubic anisotropy constant", &kc3_red)
	AnisU      = NewVectorParam("anisU", "", "Uniaxial anisotropy direction")
	AnisC1     = NewVectorParam("anisC1", "", "Cubic anisotropy direction #1")
	AnisC2     = NewVectorParam("anisC2", "", "Cubic anisotorpy directon #2")
	B_anis     = NewVectorField("B_anis", "T", "Anisotropy filed", AddAnisotropyField)
	Edens_anis = NewScalarField("Edens_anis", "J/m3", "Anisotropy energy density", AddAnisotropyEnergyDensity)
	E_anis     = NewScalarValue("E_anis", "J", "total anisotropy energy", GetAnisotropyEnergy)
	KuXX       = NewScalarParam("KuXX", "J/m3", "1st order uniaxial anisotropy constant (XX) for monodomain shape anisotropy", &kuxx_red)
	KuYY       = NewScalarParam("KuYY", "J/m3", "1st order uniaxial anisotropy constant (YY) for monodomain shape anisotropy", &kuyy_red)
	KuZZ       = NewScalarParam("KuZZ", "J/m3", "1st order uniaxial anisotropy constant (ZZ) for monodomain shape anisotropy", &kuzz_red)
	AnisUxx    = NewVectorParam("anisUxx", "", "Uniaxial anisotropy direction (XX) for monodomain shape anisotropy")
	AnisUyy    = NewVectorParam("anisUyy", "", "Uniaxial anisotropy direction (YY) for monodomain shape anisotropy")
	AnisUzz    = NewVectorParam("anisUzz", "", "Uniaxial anisotropy direction (ZZ) for monodomain shape anisotropy")
)

var (
	ku1_red  DerivedParam
	ku2_red  DerivedParam
	kc1_red  DerivedParam
	kc2_red  DerivedParam
	kc3_red  DerivedParam
	kuxx_red DerivedParam
	kuyy_red DerivedParam
	kuzz_red DerivedParam
)

var zero param // utility zero parameter

func init() {
	Ku1.addChild(&ku1_red)
	registerEnergy(GetAnisotropyEnergy, AddAnisotropyEnergyDensity)
	zero.init(1, "_zero", "", nil)

	//ku1_red = Ku1 / Msat
	ku1_red.init(SCALAR, []parent{Ku1, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, Ku1.cpuLUT(), Msat.cpuLUT())
	})
	//ku2_red = Ku2 / Msat
	ku2_red.init(SCALAR, []parent{Ku2, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, Ku2.cpuLUT(), Msat.cpuLUT())
	})

	//kc1_red = Kc1 / Msat
	kc1_red.init(SCALAR, []parent{Kc1, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, Kc1.cpuLUT(), Msat.cpuLUT())
	})
	//kc2_red = Kc2 / Msat
	kc2_red.init(SCALAR, []parent{Kc2, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, Kc2.cpuLUT(), Msat.cpuLUT())
	})
	//kc3_red = Kc3 / Msat
	kc3_red.init(SCALAR, []parent{Kc3, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, Kc3.cpuLUT(), Msat.cpuLUT())
	})
	//kuxx_red = KuXX / Msat
	kuxx_red.init(SCALAR, []parent{KuXX, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, KuXX.cpuLUT(), Msat.cpuLUT())
	})
	//kuyy_red = KuYY / Msat
	kuyy_red.init(SCALAR, []parent{KuYY, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, KuYY.cpuLUT(), Msat.cpuLUT())
	})
	//kuzz_red = KuZZ / Msat
	kuzz_red.init(SCALAR, []parent{KuZZ, Msat}, func(p *DerivedParam) {
		paramDiv(p.cpu_buf, KuZZ.cpuLUT(), Msat.cpuLUT())
	})
}

func addUniaxialAnisotropyField(dst *data.Slice) {
	if ku1_red.nonZero() || ku2_red.nonZero() {
		cuda.AddUniaxialAnisotropy(dst, M.Buffer(), ku1_red.gpuLUT1(), ku2_red.gpuLUT1(), AnisU.gpuLUT(), regions.Gpu())
	}
}

func addCubicAnisotropyField(dst *data.Slice) {
	if kc1_red.nonZero() || kc2_red.nonZero() || kc3_red.nonZero() {
		cuda.AddCubicAnisotropy(dst, M.Buffer(), kc1_red.gpuLUT1(), kc2_red.gpuLUT1(), kc3_red.gpuLUT1(), AnisC1.gpuLUT(), AnisC2.gpuLUT(), regions.Gpu())
	}
}

func addShapeAnisotropyField(dst *data.Slice) {
	if kuxx_red.nonZero() {
		cuda.AddUniaxialAnisotropy(dst, M.Buffer(), kuxx_red.gpuLUT1(), zero.gpuLUT1(), AnisUxx.gpuLUT(), regions.Gpu())
	}
	if kuyy_red.nonZero() {
		cuda.AddUniaxialAnisotropy(dst, M.Buffer(), kuyy_red.gpuLUT1(), zero.gpuLUT1(), AnisUyy.gpuLUT(), regions.Gpu())
	}
	if kuzz_red.nonZero() {
		cuda.AddUniaxialAnisotropy(dst, M.Buffer(), kuzz_red.gpuLUT1(), zero.gpuLUT1(), AnisUzz.gpuLUT(), regions.Gpu())
	}
}

// Add the anisotropy field to dst
func AddAnisotropyField(dst *data.Slice) {
	addUniaxialAnisotropyField(dst)
	addCubicAnisotropyField(dst)
	addShapeAnisotropyField(dst)
}

func AddAnisotropyEnergyDensity(dst *data.Slice) {
	haveUnixial := ku1_red.nonZero() || ku2_red.nonZero()
	haveCubic := kc1_red.nonZero() || kc2_red.nonZero() || kc3_red.nonZero()
	haveShape := kuxx_red.nonZero() || kuyy_red.nonZero() || kuzz_red.nonZero()

	if !haveUnixial && !haveCubic && !haveShape {
		return
	}

	buf := cuda.Buffer(B_anis.NComp(), B_anis.Mesh().Size())
	defer cuda.Recycle(buf)

	// unnormalized magnetization:
	Mf, r := M_full.Slice()
	if r {
		defer cuda.Recycle(Mf)
	}

	if haveUnixial {
		// 1st
		cuda.Zero(buf)
		cuda.AddUniaxialAnisotropy(buf, M.Buffer(), ku1_red.gpuLUT1(), zero.gpuLUT1(), AnisU.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./2., buf, Mf)

		// 2nd
		cuda.Zero(buf)
		cuda.AddUniaxialAnisotropy(buf, M.Buffer(), zero.gpuLUT1(), ku2_red.gpuLUT1(), AnisU.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./4., buf, Mf)
	}

	if haveCubic {
		// 1st
		cuda.Zero(buf)
		cuda.AddCubicAnisotropy(buf, M.Buffer(), kc1_red.gpuLUT1(), zero.gpuLUT1(), zero.gpuLUT1(), AnisC1.gpuLUT(), AnisC2.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./4., buf, Mf)

		// 2nd
		cuda.Zero(buf)
		cuda.AddCubicAnisotropy(buf, M.Buffer(), zero.gpuLUT1(), kc2_red.gpuLUT1(), zero.gpuLUT1(), AnisC1.gpuLUT(), AnisC2.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./6., buf, Mf)

		// 3nd
		cuda.Zero(buf)
		cuda.AddCubicAnisotropy(buf, M.Buffer(), zero.gpuLUT1(), zero.gpuLUT1(), kc3_red.gpuLUT1(), AnisC1.gpuLUT(), AnisC2.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./8., buf, Mf)
	}

	if haveShape {
		// XX
		cuda.Zero(buf)
		cuda.AddUniaxialAnisotropy(buf, M.Buffer(), kuxx_red.gpuLUT1(), zero.gpuLUT1(), AnisUxx.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./2., buf, Mf)

		// YY
		cuda.Zero(buf)
		cuda.AddUniaxialAnisotropy(buf, M.Buffer(), kuyy_red.gpuLUT1(), zero.gpuLUT1(), AnisUyy.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./2., buf, Mf)

		// ZZ
		cuda.Zero(buf)
		cuda.AddUniaxialAnisotropy(buf, M.Buffer(), kuzz_red.gpuLUT1(), zero.gpuLUT1(), AnisUzz.gpuLUT(), regions.Gpu())
		cuda.AddDotProduct(dst, -1./2., buf, Mf)
	}
}

// Returns anisotropy energy in joules.
func GetAnisotropyEnergy() float64 {
	buf := cuda.Buffer(1, Edens_anis.Mesh().Size())
	defer cuda.Recycle(buf)

	cuda.Zero(buf)
	AddAnisotropyEnergyDensity(buf)
	return cellVolume() * float64(cuda.Sum(buf))
}
