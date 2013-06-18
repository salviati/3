package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var addexchange_code cu.Function

type addexchange_args struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_aLUT2d  unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_wx      float32
	arg_wy      float32
	arg_wz      float32
	arg_N0      int
	arg_N1      int
	arg_N2      int
	argptr      [14]unsafe.Pointer
}

// Wrapper for addexchange CUDA kernel, asynchronous.
func k_addexchange_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, aLUT2d unsafe.Pointer, regions unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if addexchange_code == 0 {
		addexchange_code = fatbinLoad(addexchange_map, "addexchange")
	}

	var a addexchange_args

	a.arg_Bx = Bx
	a.argptr[0] = unsafe.Pointer(&a.arg_Bx)
	a.arg_By = By
	a.argptr[1] = unsafe.Pointer(&a.arg_By)
	a.arg_Bz = Bz
	a.argptr[2] = unsafe.Pointer(&a.arg_Bz)
	a.arg_mx = mx
	a.argptr[3] = unsafe.Pointer(&a.arg_mx)
	a.arg_my = my
	a.argptr[4] = unsafe.Pointer(&a.arg_my)
	a.arg_mz = mz
	a.argptr[5] = unsafe.Pointer(&a.arg_mz)
	a.arg_aLUT2d = aLUT2d
	a.argptr[6] = unsafe.Pointer(&a.arg_aLUT2d)
	a.arg_regions = regions
	a.argptr[7] = unsafe.Pointer(&a.arg_regions)
	a.arg_wx = wx
	a.argptr[8] = unsafe.Pointer(&a.arg_wx)
	a.arg_wy = wy
	a.argptr[9] = unsafe.Pointer(&a.arg_wy)
	a.arg_wz = wz
	a.argptr[10] = unsafe.Pointer(&a.arg_wz)
	a.arg_N0 = N0
	a.argptr[11] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[12] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[13] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(addexchange_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for addexchange CUDA kernel, synchronized.
func k_addexchange(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, aLUT2d unsafe.Pointer, regions unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_addexchange_async(Bx, By, Bz, mx, my, mz, aLUT2d, regions, wx, wy, wz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var addexchange_map = map[int]string{0: "",
	20: addexchange_ptx_20,
	30: addexchange_ptx_30,
	35: addexchange_ptx_35}

const (
	addexchange_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry addexchange(
	.param .u64 addexchange_param_0,
	.param .u64 addexchange_param_1,
	.param .u64 addexchange_param_2,
	.param .u64 addexchange_param_3,
	.param .u64 addexchange_param_4,
	.param .u64 addexchange_param_5,
	.param .u64 addexchange_param_6,
	.param .u64 addexchange_param_7,
	.param .f32 addexchange_param_8,
	.param .f32 addexchange_param_9,
	.param .f32 addexchange_param_10,
	.param .u32 addexchange_param_11,
	.param .u32 addexchange_param_12,
	.param .u32 addexchange_param_13
)
{
	.reg .pred 	%p<14>;
	.reg .s16 	%rc<8>;
	.reg .s32 	%r<210>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<99>;


	ld.param.u64 	%rd4, [addexchange_param_0];
	ld.param.u64 	%rd5, [addexchange_param_1];
	ld.param.u64 	%rd6, [addexchange_param_2];
	ld.param.u64 	%rd7, [addexchange_param_3];
	ld.param.u64 	%rd8, [addexchange_param_4];
	ld.param.u64 	%rd9, [addexchange_param_5];
	ld.param.u64 	%rd10, [addexchange_param_6];
	ld.param.u64 	%rd11, [addexchange_param_7];
	ld.param.f32 	%f46, [addexchange_param_8];
	ld.param.f32 	%f47, [addexchange_param_9];
	ld.param.f32 	%f48, [addexchange_param_10];
	ld.param.u32 	%r56, [addexchange_param_11];
	ld.param.u32 	%r57, [addexchange_param_12];
	ld.param.u32 	%r58, [addexchange_param_13];
	.loc 2 16 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 17 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 19 1
	setp.lt.s32 	%p1, %r8, %r58;
	setp.lt.s32 	%p2, %r4, %r57;
	and.pred  	%p3, %p2, %p1;
	.loc 2 23 1
	setp.gt.s32 	%p4, %r56, 0;
	.loc 2 19 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_23;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 36 1
	add.s32 	%r60, %r8, -1;
	mov.u32 	%r203, 0;
	.loc 3 238 5
	max.s32 	%r61, %r60, %r203;
	.loc 2 42 1
	add.s32 	%r62, %r58, -1;
	add.s32 	%r63, %r8, 1;
	.loc 3 210 5
	min.s32 	%r64, %r63, %r62;
	.loc 2 48 1
	add.s32 	%r65, %r4, -1;
	.loc 3 238 5
	max.s32 	%r66, %r65, %r203;
	.loc 2 54 1
	add.s32 	%r67, %r57, -1;
	add.s32 	%r68, %r4, 1;
	.loc 3 210 5
	min.s32 	%r69, %r68, %r67;
	.loc 2 23 1
	mad.lo.s32 	%r202, %r69, %r58, %r8;
	mad.lo.s32 	%r201, %r66, %r58, %r8;
	mad.lo.s32 	%r200, %r58, %r4, %r64;
	mad.lo.s32 	%r199, %r58, %r4, %r61;
	mad.lo.s32 	%r198, %r58, %r4, %r8;
	cvta.to.global.u64 	%rd22, %rd4;
	cvta.to.global.u64 	%rd23, %rd5;
	cvta.to.global.u64 	%rd24, %rd6;

BB0_2:
	cvta.to.global.u64 	%rd12, %rd11;
	cvta.to.global.u64 	%rd13, %rd7;
	.loc 2 27 1
	cvt.s64.s32 	%rd14, %r198;
	mul.wide.s32 	%rd15, %r198, 4;
	add.s64 	%rd16, %rd13, %rd15;
	ld.global.f32 	%f1, [%rd16];
	cvta.to.global.u64 	%rd17, %rd8;
	.loc 2 27 1
	add.s64 	%rd18, %rd17, %rd15;
	ld.global.f32 	%f2, [%rd18];
	cvta.to.global.u64 	%rd19, %rd9;
	.loc 2 27 1
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f3, [%rd20];
	.loc 2 28 1
	add.s64 	%rd21, %rd12, %rd14;
	.loc 2 29 1
	add.s64 	%rd1, %rd22, %rd15;
	ld.global.f32 	%f4, [%rd1];
	add.s64 	%rd2, %rd23, %rd15;
	ld.global.f32 	%f5, [%rd2];
	add.s64 	%rd3, %rd24, %rd15;
	ld.global.f32 	%f6, [%rd3];
	.loc 2 37 1
	cvt.s64.s32 	%rd25, %r199;
	mul.wide.s32 	%rd26, %r199, 4;
	add.s64 	%rd27, %rd13, %rd26;
	ld.global.f32 	%f7, [%rd27];
	add.s64 	%rd28, %rd17, %rd26;
	ld.global.f32 	%f8, [%rd28];
	add.s64 	%rd29, %rd19, %rd26;
	ld.global.f32 	%f9, [%rd29];
	.loc 2 38 1
	add.s64 	%rd30, %rd12, %rd25;
	ld.global.u8 	%rc2, [%rd30];
	.loc 2 28 1
	ld.global.u8 	%rc1, [%rd21];
	.loc 2 38 1
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc2;
	setp.gt.s16 	%p6, %temp1, %temp2;
	}
	cvt.s32.s8 	%r20, %rc2;
	cvt.s32.s8 	%r21, %rc1;
	@%p6 bra 	BB0_4;

	add.s32 	%r83, %r21, 1;
	mul.lo.s32 	%r84, %r83, %r21;
	shr.u32 	%r85, %r84, 31;
	mad.lo.s32 	%r86, %r83, %r21, %r85;
	shr.s32 	%r87, %r86, 1;
	add.s32 	%r204, %r20, %r87;
	bra.uni 	BB0_5;

BB0_4:
	.loc 2 38 1
	add.s32 	%r88, %r20, 1;
	mul.lo.s32 	%r89, %r88, %r20;
	shr.u32 	%r90, %r89, 31;
	mad.lo.s32 	%r91, %r88, %r20, %r90;
	shr.s32 	%r92, %r91, 1;
	add.s32 	%r204, %r92, %r21;

BB0_5:
	cvta.to.global.u64 	%rd32, %rd10;
	.loc 2 38 1
	mul.wide.s32 	%rd33, %r204, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.f32 	%f49, [%rd34];
	.loc 2 39 1
	mul.f32 	%f50, %f49, %f48;
	sub.f32 	%f51, %f7, %f1;
	sub.f32 	%f52, %f8, %f2;
	sub.f32 	%f53, %f9, %f3;
	.loc 2 39 1
	fma.rn.f32 	%f10, %f50, %f51, %f4;
	fma.rn.f32 	%f11, %f50, %f52, %f5;
	fma.rn.f32 	%f12, %f50, %f53, %f6;
	.loc 2 43 1
	cvt.s64.s32 	%rd36, %r200;
	mul.wide.s32 	%rd37, %r200, 4;
	add.s64 	%rd38, %rd13, %rd37;
	ld.global.f32 	%f13, [%rd38];
	add.s64 	%rd40, %rd17, %rd37;
	ld.global.f32 	%f14, [%rd40];
	add.s64 	%rd42, %rd19, %rd37;
	ld.global.f32 	%f15, [%rd42];
	.loc 2 44 1
	add.s64 	%rd43, %rd12, %rd36;
	ld.global.u8 	%rc3, [%rd43];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc3;
	setp.gt.s16 	%p7, %temp1, %temp2;
	}
	cvt.s32.s8 	%r26, %rc3;
	@%p7 bra 	BB0_7;

	add.s32 	%r98, %r21, 1;
	mul.lo.s32 	%r99, %r98, %r21;
	shr.u32 	%r100, %r99, 31;
	mad.lo.s32 	%r101, %r98, %r21, %r100;
	shr.s32 	%r102, %r101, 1;
	add.s32 	%r205, %r26, %r102;
	bra.uni 	BB0_8;

BB0_7:
	.loc 2 44 1
	add.s32 	%r103, %r26, 1;
	mul.lo.s32 	%r104, %r103, %r26;
	shr.u32 	%r105, %r104, 31;
	mad.lo.s32 	%r106, %r103, %r26, %r105;
	shr.s32 	%r107, %r106, 1;
	add.s32 	%r205, %r107, %r21;

BB0_8:
	mul.wide.s32 	%rd46, %r205, 4;
	add.s64 	%rd47, %rd32, %rd46;
	ld.global.f32 	%f54, [%rd47];
	.loc 2 45 1
	mul.f32 	%f55, %f54, %f48;
	sub.f32 	%f56, %f13, %f1;
	sub.f32 	%f57, %f14, %f2;
	sub.f32 	%f58, %f15, %f3;
	.loc 2 45 1
	fma.rn.f32 	%f16, %f55, %f56, %f10;
	fma.rn.f32 	%f17, %f55, %f57, %f11;
	fma.rn.f32 	%f18, %f55, %f58, %f12;
	.loc 2 49 1
	cvt.s64.s32 	%rd49, %r201;
	mul.wide.s32 	%rd50, %r201, 4;
	add.s64 	%rd51, %rd13, %rd50;
	ld.global.f32 	%f19, [%rd51];
	add.s64 	%rd53, %rd17, %rd50;
	ld.global.f32 	%f20, [%rd53];
	add.s64 	%rd55, %rd19, %rd50;
	ld.global.f32 	%f21, [%rd55];
	.loc 2 50 1
	add.s64 	%rd56, %rd12, %rd49;
	ld.global.u8 	%rc4, [%rd56];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc4;
	setp.gt.s16 	%p8, %temp1, %temp2;
	}
	cvt.s32.s8 	%r31, %rc4;
	@%p8 bra 	BB0_10;

	add.s32 	%r113, %r21, 1;
	mul.lo.s32 	%r114, %r113, %r21;
	shr.u32 	%r115, %r114, 31;
	mad.lo.s32 	%r116, %r113, %r21, %r115;
	shr.s32 	%r117, %r116, 1;
	add.s32 	%r206, %r31, %r117;
	bra.uni 	BB0_11;

BB0_10:
	.loc 2 50 1
	add.s32 	%r118, %r31, 1;
	mul.lo.s32 	%r119, %r118, %r31;
	shr.u32 	%r120, %r119, 31;
	mad.lo.s32 	%r121, %r118, %r31, %r120;
	shr.s32 	%r122, %r121, 1;
	add.s32 	%r206, %r122, %r21;

BB0_11:
	mul.wide.s32 	%rd59, %r206, 4;
	add.s64 	%rd60, %rd32, %rd59;
	ld.global.f32 	%f59, [%rd60];
	.loc 2 51 1
	mul.f32 	%f60, %f59, %f47;
	sub.f32 	%f61, %f19, %f1;
	sub.f32 	%f62, %f20, %f2;
	sub.f32 	%f63, %f21, %f3;
	.loc 2 51 1
	fma.rn.f32 	%f22, %f60, %f61, %f16;
	fma.rn.f32 	%f23, %f60, %f62, %f17;
	fma.rn.f32 	%f24, %f60, %f63, %f18;
	.loc 2 55 1
	cvt.s64.s32 	%rd62, %r202;
	mul.wide.s32 	%rd63, %r202, 4;
	add.s64 	%rd64, %rd13, %rd63;
	ld.global.f32 	%f25, [%rd64];
	add.s64 	%rd66, %rd17, %rd63;
	ld.global.f32 	%f26, [%rd66];
	add.s64 	%rd68, %rd19, %rd63;
	ld.global.f32 	%f27, [%rd68];
	.loc 2 56 1
	add.s64 	%rd69, %rd12, %rd62;
	ld.global.u8 	%rc5, [%rd69];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc5;
	setp.gt.s16 	%p9, %temp1, %temp2;
	}
	cvt.s32.s8 	%r36, %rc5;
	@%p9 bra 	BB0_13;

	add.s32 	%r128, %r21, 1;
	mul.lo.s32 	%r129, %r128, %r21;
	shr.u32 	%r130, %r129, 31;
	mad.lo.s32 	%r131, %r128, %r21, %r130;
	shr.s32 	%r132, %r131, 1;
	add.s32 	%r207, %r36, %r132;
	bra.uni 	BB0_14;

BB0_13:
	.loc 2 56 1
	add.s32 	%r133, %r36, 1;
	mul.lo.s32 	%r134, %r133, %r36;
	shr.u32 	%r135, %r134, 31;
	mad.lo.s32 	%r136, %r133, %r36, %r135;
	shr.s32 	%r137, %r136, 1;
	add.s32 	%r207, %r137, %r21;

BB0_14:
	mul.wide.s32 	%rd71, %r207, 4;
	add.s64 	%rd72, %rd32, %rd71;
	ld.global.f32 	%f64, [%rd72];
	.loc 2 57 1
	mul.f32 	%f65, %f64, %f47;
	sub.f32 	%f66, %f25, %f1;
	sub.f32 	%f67, %f26, %f2;
	sub.f32 	%f68, %f27, %f3;
	.loc 2 57 1
	fma.rn.f32 	%f79, %f65, %f66, %f22;
	fma.rn.f32 	%f80, %f65, %f67, %f23;
	fma.rn.f32 	%f81, %f65, %f68, %f24;
	setp.eq.s32 	%p10, %r56, 1;
	.loc 2 60 1
	@%p10 bra 	BB0_22;

	.loc 2 62 1
	add.s32 	%r139, %r203, -1;
	mov.u32 	%r140, 0;
	.loc 3 238 5
	max.s32 	%r141, %r139, %r140;
	.loc 2 62 1
	mad.lo.s32 	%r146, %r141, %r57, %r4;
	mad.lo.s32 	%r151, %r146, %r58, %r8;
	.loc 2 63 1
	cvt.s64.s32 	%rd74, %r151;
	mul.wide.s32 	%rd76, %r151, 4;
	add.s64 	%rd77, %rd13, %rd76;
	ld.global.f32 	%f31, [%rd77];
	add.s64 	%rd79, %rd17, %rd76;
	ld.global.f32 	%f32, [%rd79];
	add.s64 	%rd81, %rd19, %rd76;
	ld.global.f32 	%f33, [%rd81];
	.loc 2 64 1
	add.s64 	%rd82, %rd12, %rd74;
	ld.global.u8 	%rc6, [%rd82];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc6;
	setp.gt.s16 	%p11, %temp1, %temp2;
	}
	cvt.s32.s8 	%r41, %rc6;
	@%p11 bra 	BB0_17;

	add.s32 	%r156, %r21, 1;
	mul.lo.s32 	%r157, %r156, %r21;
	shr.u32 	%r158, %r157, 31;
	mad.lo.s32 	%r159, %r156, %r21, %r158;
	shr.s32 	%r160, %r159, 1;
	add.s32 	%r208, %r41, %r160;
	bra.uni 	BB0_18;

BB0_17:
	.loc 2 64 1
	add.s32 	%r161, %r41, 1;
	mul.lo.s32 	%r162, %r161, %r41;
	shr.u32 	%r163, %r162, 31;
	mad.lo.s32 	%r164, %r161, %r41, %r163;
	shr.s32 	%r165, %r164, 1;
	add.s32 	%r208, %r165, %r21;

BB0_18:
	mul.wide.s32 	%rd85, %r208, 4;
	add.s64 	%rd86, %rd32, %rd85;
	ld.global.f32 	%f69, [%rd86];
	.loc 2 65 1
	mul.f32 	%f70, %f69, %f46;
	sub.f32 	%f71, %f31, %f1;
	sub.f32 	%f72, %f32, %f2;
	sub.f32 	%f73, %f33, %f3;
	.loc 2 65 1
	fma.rn.f32 	%f34, %f70, %f71, %f79;
	fma.rn.f32 	%f35, %f70, %f72, %f80;
	fma.rn.f32 	%f36, %f70, %f73, %f81;
	.loc 2 68 1
	add.s32 	%r167, %r56, -1;
	add.s32 	%r168, %r203, 1;
	.loc 3 210 5
	min.s32 	%r169, %r168, %r167;
	.loc 2 68 1
	mad.lo.s32 	%r174, %r169, %r57, %r4;
	mad.lo.s32 	%r179, %r174, %r58, %r8;
	.loc 2 69 1
	cvt.s64.s32 	%rd87, %r179;
	mul.wide.s32 	%rd89, %r179, 4;
	add.s64 	%rd90, %rd13, %rd89;
	ld.global.f32 	%f37, [%rd90];
	add.s64 	%rd92, %rd17, %rd89;
	ld.global.f32 	%f38, [%rd92];
	add.s64 	%rd94, %rd19, %rd89;
	ld.global.f32 	%f39, [%rd94];
	.loc 2 70 1
	add.s64 	%rd95, %rd12, %rd87;
	ld.global.u8 	%rc7, [%rd95];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc7;
	setp.gt.s16 	%p12, %temp1, %temp2;
	}
	cvt.s32.s8 	%r46, %rc7;
	@%p12 bra 	BB0_20;

	add.s32 	%r184, %r21, 1;
	mul.lo.s32 	%r185, %r184, %r21;
	shr.u32 	%r186, %r185, 31;
	mad.lo.s32 	%r187, %r184, %r21, %r186;
	shr.s32 	%r188, %r187, 1;
	add.s32 	%r209, %r46, %r188;
	bra.uni 	BB0_21;

BB0_20:
	.loc 2 70 1
	add.s32 	%r189, %r46, 1;
	mul.lo.s32 	%r190, %r189, %r46;
	shr.u32 	%r191, %r190, 31;
	mad.lo.s32 	%r192, %r189, %r46, %r191;
	shr.s32 	%r193, %r192, 1;
	add.s32 	%r209, %r193, %r21;

BB0_21:
	mul.wide.s32 	%rd97, %r209, 4;
	add.s64 	%rd98, %rd32, %rd97;
	ld.global.f32 	%f74, [%rd98];
	.loc 2 71 1
	mul.f32 	%f75, %f74, %f46;
	sub.f32 	%f76, %f37, %f1;
	sub.f32 	%f77, %f38, %f2;
	sub.f32 	%f78, %f39, %f3;
	.loc 2 71 1
	fma.rn.f32 	%f79, %f75, %f76, %f34;
	fma.rn.f32 	%f80, %f75, %f77, %f35;
	fma.rn.f32 	%f81, %f75, %f78, %f36;

BB0_22:
	.loc 2 74 1
	st.global.f32 	[%rd1], %f79;
	.loc 2 75 1
	st.global.f32 	[%rd2], %f80;
	.loc 2 76 1
	st.global.f32 	[%rd3], %f81;
	.loc 2 23 1
	mad.lo.s32 	%r202, %r58, %r57, %r202;
	mad.lo.s32 	%r201, %r58, %r57, %r201;
	mad.lo.s32 	%r200, %r58, %r57, %r200;
	mad.lo.s32 	%r199, %r58, %r57, %r199;
	mad.lo.s32 	%r198, %r58, %r57, %r198;
	.loc 2 23 18
	add.s32 	%r203, %r203, 1;
	.loc 2 23 1
	setp.lt.s32 	%p13, %r203, %r56;
	@%p13 bra 	BB0_2;

BB0_23:
	.loc 2 78 2
	ret;
}


`
	addexchange_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry addexchange(
	.param .u64 addexchange_param_0,
	.param .u64 addexchange_param_1,
	.param .u64 addexchange_param_2,
	.param .u64 addexchange_param_3,
	.param .u64 addexchange_param_4,
	.param .u64 addexchange_param_5,
	.param .u64 addexchange_param_6,
	.param .u64 addexchange_param_7,
	.param .f32 addexchange_param_8,
	.param .f32 addexchange_param_9,
	.param .f32 addexchange_param_10,
	.param .u32 addexchange_param_11,
	.param .u32 addexchange_param_12,
	.param .u32 addexchange_param_13
)
{
	.reg .pred 	%p<14>;
	.reg .s16 	%rc<8>;
	.reg .s32 	%r<195>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<99>;


	ld.param.u64 	%rd4, [addexchange_param_0];
	ld.param.u64 	%rd5, [addexchange_param_1];
	ld.param.u64 	%rd6, [addexchange_param_2];
	ld.param.u64 	%rd7, [addexchange_param_3];
	ld.param.u64 	%rd8, [addexchange_param_4];
	ld.param.u64 	%rd9, [addexchange_param_5];
	ld.param.u64 	%rd10, [addexchange_param_6];
	ld.param.u64 	%rd11, [addexchange_param_7];
	ld.param.f32 	%f46, [addexchange_param_8];
	ld.param.f32 	%f47, [addexchange_param_9];
	ld.param.f32 	%f48, [addexchange_param_10];
	ld.param.u32 	%r57, [addexchange_param_11];
	ld.param.u32 	%r58, [addexchange_param_12];
	ld.param.u32 	%r59, [addexchange_param_13];
	.loc 2 16 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 17 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 19 1
	setp.lt.s32 	%p1, %r8, %r59;
	setp.lt.s32 	%p2, %r4, %r58;
	and.pred  	%p3, %p2, %p1;
	.loc 2 23 1
	setp.gt.s32 	%p4, %r57, 0;
	.loc 2 19 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_23;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 36 1
	add.s32 	%r61, %r8, -1;
	mov.u32 	%r188, 0;
	.loc 3 238 5
	max.s32 	%r62, %r61, %r188;
	.loc 2 42 1
	add.s32 	%r63, %r59, -1;
	add.s32 	%r64, %r8, 1;
	.loc 3 210 5
	min.s32 	%r65, %r64, %r63;
	.loc 2 48 1
	add.s32 	%r66, %r4, -1;
	.loc 3 238 5
	max.s32 	%r67, %r66, %r188;
	.loc 2 54 1
	add.s32 	%r68, %r58, -1;
	add.s32 	%r69, %r4, 1;
	.loc 3 210 5
	min.s32 	%r70, %r69, %r68;
	.loc 2 23 1
	mad.lo.s32 	%r187, %r70, %r59, %r8;
	mul.lo.s32 	%r10, %r59, %r58;
	mad.lo.s32 	%r186, %r67, %r59, %r8;
	mad.lo.s32 	%r185, %r59, %r4, %r65;
	mad.lo.s32 	%r184, %r59, %r4, %r62;
	mad.lo.s32 	%r183, %r59, %r4, %r8;
	cvta.to.global.u64 	%rd22, %rd4;
	cvta.to.global.u64 	%rd23, %rd5;
	cvta.to.global.u64 	%rd24, %rd6;

BB0_2:
	cvta.to.global.u64 	%rd12, %rd11;
	cvta.to.global.u64 	%rd13, %rd7;
	.loc 2 27 1
	cvt.s64.s32 	%rd14, %r183;
	mul.wide.s32 	%rd15, %r183, 4;
	add.s64 	%rd16, %rd13, %rd15;
	ld.global.f32 	%f1, [%rd16];
	cvta.to.global.u64 	%rd17, %rd8;
	.loc 2 27 1
	add.s64 	%rd18, %rd17, %rd15;
	ld.global.f32 	%f2, [%rd18];
	cvta.to.global.u64 	%rd19, %rd9;
	.loc 2 27 1
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f3, [%rd20];
	.loc 2 28 1
	add.s64 	%rd21, %rd12, %rd14;
	.loc 2 29 1
	add.s64 	%rd1, %rd22, %rd15;
	ld.global.f32 	%f4, [%rd1];
	add.s64 	%rd2, %rd23, %rd15;
	ld.global.f32 	%f5, [%rd2];
	add.s64 	%rd3, %rd24, %rd15;
	ld.global.f32 	%f6, [%rd3];
	.loc 2 37 1
	cvt.s64.s32 	%rd25, %r184;
	mul.wide.s32 	%rd26, %r184, 4;
	add.s64 	%rd27, %rd13, %rd26;
	ld.global.f32 	%f7, [%rd27];
	add.s64 	%rd28, %rd17, %rd26;
	ld.global.f32 	%f8, [%rd28];
	add.s64 	%rd29, %rd19, %rd26;
	ld.global.f32 	%f9, [%rd29];
	.loc 2 38 1
	add.s64 	%rd30, %rd12, %rd25;
	ld.global.u8 	%rc2, [%rd30];
	.loc 2 28 1
	ld.global.u8 	%rc1, [%rd21];
	.loc 2 38 1
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc2;
	setp.gt.s16 	%p6, %temp1, %temp2;
	}
	cvt.s32.s8 	%r21, %rc2;
	cvt.s32.s8 	%r22, %rc1;
	@%p6 bra 	BB0_4;

	add.s32 	%r84, %r22, 1;
	mul.lo.s32 	%r85, %r84, %r22;
	shr.u32 	%r86, %r85, 31;
	mad.lo.s32 	%r87, %r84, %r22, %r86;
	shr.s32 	%r88, %r87, 1;
	add.s32 	%r189, %r21, %r88;
	bra.uni 	BB0_5;

BB0_4:
	.loc 2 38 1
	add.s32 	%r89, %r21, 1;
	mul.lo.s32 	%r90, %r89, %r21;
	shr.u32 	%r91, %r90, 31;
	mad.lo.s32 	%r92, %r89, %r21, %r91;
	shr.s32 	%r93, %r92, 1;
	add.s32 	%r189, %r93, %r22;

BB0_5:
	cvta.to.global.u64 	%rd32, %rd10;
	.loc 2 38 1
	mul.wide.s32 	%rd33, %r189, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.f32 	%f49, [%rd34];
	.loc 2 39 1
	mul.f32 	%f50, %f49, %f48;
	sub.f32 	%f51, %f7, %f1;
	sub.f32 	%f52, %f8, %f2;
	sub.f32 	%f53, %f9, %f3;
	.loc 2 39 1
	fma.rn.f32 	%f10, %f50, %f51, %f4;
	fma.rn.f32 	%f11, %f50, %f52, %f5;
	fma.rn.f32 	%f12, %f50, %f53, %f6;
	.loc 2 43 1
	cvt.s64.s32 	%rd36, %r185;
	mul.wide.s32 	%rd37, %r185, 4;
	add.s64 	%rd38, %rd13, %rd37;
	ld.global.f32 	%f13, [%rd38];
	add.s64 	%rd40, %rd17, %rd37;
	ld.global.f32 	%f14, [%rd40];
	add.s64 	%rd42, %rd19, %rd37;
	ld.global.f32 	%f15, [%rd42];
	.loc 2 44 1
	add.s64 	%rd43, %rd12, %rd36;
	ld.global.u8 	%rc3, [%rd43];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc3;
	setp.gt.s16 	%p7, %temp1, %temp2;
	}
	cvt.s32.s8 	%r27, %rc3;
	@%p7 bra 	BB0_7;

	add.s32 	%r99, %r22, 1;
	mul.lo.s32 	%r100, %r99, %r22;
	shr.u32 	%r101, %r100, 31;
	mad.lo.s32 	%r102, %r99, %r22, %r101;
	shr.s32 	%r103, %r102, 1;
	add.s32 	%r190, %r27, %r103;
	bra.uni 	BB0_8;

BB0_7:
	.loc 2 44 1
	add.s32 	%r104, %r27, 1;
	mul.lo.s32 	%r105, %r104, %r27;
	shr.u32 	%r106, %r105, 31;
	mad.lo.s32 	%r107, %r104, %r27, %r106;
	shr.s32 	%r108, %r107, 1;
	add.s32 	%r190, %r108, %r22;

BB0_8:
	mul.wide.s32 	%rd46, %r190, 4;
	add.s64 	%rd47, %rd32, %rd46;
	ld.global.f32 	%f54, [%rd47];
	.loc 2 45 1
	mul.f32 	%f55, %f54, %f48;
	sub.f32 	%f56, %f13, %f1;
	sub.f32 	%f57, %f14, %f2;
	sub.f32 	%f58, %f15, %f3;
	.loc 2 45 1
	fma.rn.f32 	%f16, %f55, %f56, %f10;
	fma.rn.f32 	%f17, %f55, %f57, %f11;
	fma.rn.f32 	%f18, %f55, %f58, %f12;
	.loc 2 49 1
	cvt.s64.s32 	%rd49, %r186;
	mul.wide.s32 	%rd50, %r186, 4;
	add.s64 	%rd51, %rd13, %rd50;
	ld.global.f32 	%f19, [%rd51];
	add.s64 	%rd53, %rd17, %rd50;
	ld.global.f32 	%f20, [%rd53];
	add.s64 	%rd55, %rd19, %rd50;
	ld.global.f32 	%f21, [%rd55];
	.loc 2 50 1
	add.s64 	%rd56, %rd12, %rd49;
	ld.global.u8 	%rc4, [%rd56];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc4;
	setp.gt.s16 	%p8, %temp1, %temp2;
	}
	cvt.s32.s8 	%r32, %rc4;
	@%p8 bra 	BB0_10;

	add.s32 	%r114, %r22, 1;
	mul.lo.s32 	%r115, %r114, %r22;
	shr.u32 	%r116, %r115, 31;
	mad.lo.s32 	%r117, %r114, %r22, %r116;
	shr.s32 	%r118, %r117, 1;
	add.s32 	%r191, %r32, %r118;
	bra.uni 	BB0_11;

BB0_10:
	.loc 2 50 1
	add.s32 	%r119, %r32, 1;
	mul.lo.s32 	%r120, %r119, %r32;
	shr.u32 	%r121, %r120, 31;
	mad.lo.s32 	%r122, %r119, %r32, %r121;
	shr.s32 	%r123, %r122, 1;
	add.s32 	%r191, %r123, %r22;

BB0_11:
	mul.wide.s32 	%rd59, %r191, 4;
	add.s64 	%rd60, %rd32, %rd59;
	ld.global.f32 	%f59, [%rd60];
	.loc 2 51 1
	mul.f32 	%f60, %f59, %f47;
	sub.f32 	%f61, %f19, %f1;
	sub.f32 	%f62, %f20, %f2;
	sub.f32 	%f63, %f21, %f3;
	.loc 2 51 1
	fma.rn.f32 	%f22, %f60, %f61, %f16;
	fma.rn.f32 	%f23, %f60, %f62, %f17;
	fma.rn.f32 	%f24, %f60, %f63, %f18;
	.loc 2 55 1
	cvt.s64.s32 	%rd62, %r187;
	mul.wide.s32 	%rd63, %r187, 4;
	add.s64 	%rd64, %rd13, %rd63;
	ld.global.f32 	%f25, [%rd64];
	add.s64 	%rd66, %rd17, %rd63;
	ld.global.f32 	%f26, [%rd66];
	add.s64 	%rd68, %rd19, %rd63;
	ld.global.f32 	%f27, [%rd68];
	.loc 2 56 1
	add.s64 	%rd69, %rd12, %rd62;
	ld.global.u8 	%rc5, [%rd69];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc5;
	setp.gt.s16 	%p9, %temp1, %temp2;
	}
	cvt.s32.s8 	%r37, %rc5;
	@%p9 bra 	BB0_13;

	add.s32 	%r129, %r22, 1;
	mul.lo.s32 	%r130, %r129, %r22;
	shr.u32 	%r131, %r130, 31;
	mad.lo.s32 	%r132, %r129, %r22, %r131;
	shr.s32 	%r133, %r132, 1;
	add.s32 	%r192, %r37, %r133;
	bra.uni 	BB0_14;

BB0_13:
	.loc 2 56 1
	add.s32 	%r134, %r37, 1;
	mul.lo.s32 	%r135, %r134, %r37;
	shr.u32 	%r136, %r135, 31;
	mad.lo.s32 	%r137, %r134, %r37, %r136;
	shr.s32 	%r138, %r137, 1;
	add.s32 	%r192, %r138, %r22;

BB0_14:
	mul.wide.s32 	%rd71, %r192, 4;
	add.s64 	%rd72, %rd32, %rd71;
	ld.global.f32 	%f64, [%rd72];
	.loc 2 57 1
	mul.f32 	%f65, %f64, %f47;
	sub.f32 	%f66, %f25, %f1;
	sub.f32 	%f67, %f26, %f2;
	sub.f32 	%f68, %f27, %f3;
	.loc 2 57 1
	fma.rn.f32 	%f79, %f65, %f66, %f22;
	fma.rn.f32 	%f80, %f65, %f67, %f23;
	fma.rn.f32 	%f81, %f65, %f68, %f24;
	setp.eq.s32 	%p10, %r57, 1;
	.loc 2 60 1
	@%p10 bra 	BB0_22;

	.loc 2 62 1
	add.s32 	%r140, %r188, -1;
	mov.u32 	%r141, 0;
	.loc 3 238 5
	max.s32 	%r142, %r140, %r141;
	.loc 2 62 1
	mad.lo.s32 	%r143, %r142, %r58, %r4;
	mad.lo.s32 	%r144, %r143, %r59, %r8;
	.loc 2 63 1
	cvt.s64.s32 	%rd74, %r144;
	mul.wide.s32 	%rd76, %r144, 4;
	add.s64 	%rd77, %rd13, %rd76;
	ld.global.f32 	%f31, [%rd77];
	add.s64 	%rd79, %rd17, %rd76;
	ld.global.f32 	%f32, [%rd79];
	add.s64 	%rd81, %rd19, %rd76;
	ld.global.f32 	%f33, [%rd81];
	.loc 2 64 1
	add.s64 	%rd82, %rd12, %rd74;
	ld.global.u8 	%rc6, [%rd82];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc6;
	setp.gt.s16 	%p11, %temp1, %temp2;
	}
	cvt.s32.s8 	%r42, %rc6;
	@%p11 bra 	BB0_17;

	add.s32 	%r149, %r22, 1;
	mul.lo.s32 	%r150, %r149, %r22;
	shr.u32 	%r151, %r150, 31;
	mad.lo.s32 	%r152, %r149, %r22, %r151;
	shr.s32 	%r153, %r152, 1;
	add.s32 	%r193, %r42, %r153;
	bra.uni 	BB0_18;

BB0_17:
	.loc 2 64 1
	add.s32 	%r154, %r42, 1;
	mul.lo.s32 	%r155, %r154, %r42;
	shr.u32 	%r156, %r155, 31;
	mad.lo.s32 	%r157, %r154, %r42, %r156;
	shr.s32 	%r158, %r157, 1;
	add.s32 	%r193, %r158, %r22;

BB0_18:
	mul.wide.s32 	%rd85, %r193, 4;
	add.s64 	%rd86, %rd32, %rd85;
	ld.global.f32 	%f69, [%rd86];
	.loc 2 65 1
	mul.f32 	%f70, %f69, %f46;
	sub.f32 	%f71, %f31, %f1;
	sub.f32 	%f72, %f32, %f2;
	sub.f32 	%f73, %f33, %f3;
	.loc 2 65 1
	fma.rn.f32 	%f34, %f70, %f71, %f79;
	fma.rn.f32 	%f35, %f70, %f72, %f80;
	fma.rn.f32 	%f36, %f70, %f73, %f81;
	.loc 2 68 1
	add.s32 	%r160, %r57, -1;
	add.s32 	%r161, %r188, 1;
	.loc 3 210 5
	min.s32 	%r162, %r161, %r160;
	.loc 2 68 1
	mad.lo.s32 	%r163, %r162, %r58, %r4;
	mad.lo.s32 	%r164, %r163, %r59, %r8;
	.loc 2 69 1
	cvt.s64.s32 	%rd87, %r164;
	mul.wide.s32 	%rd89, %r164, 4;
	add.s64 	%rd90, %rd13, %rd89;
	ld.global.f32 	%f37, [%rd90];
	add.s64 	%rd92, %rd17, %rd89;
	ld.global.f32 	%f38, [%rd92];
	add.s64 	%rd94, %rd19, %rd89;
	ld.global.f32 	%f39, [%rd94];
	.loc 2 70 1
	add.s64 	%rd95, %rd12, %rd87;
	ld.global.u8 	%rc7, [%rd95];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc7;
	setp.gt.s16 	%p12, %temp1, %temp2;
	}
	cvt.s32.s8 	%r47, %rc7;
	@%p12 bra 	BB0_20;

	add.s32 	%r169, %r22, 1;
	mul.lo.s32 	%r170, %r169, %r22;
	shr.u32 	%r171, %r170, 31;
	mad.lo.s32 	%r172, %r169, %r22, %r171;
	shr.s32 	%r173, %r172, 1;
	add.s32 	%r194, %r47, %r173;
	bra.uni 	BB0_21;

BB0_20:
	.loc 2 70 1
	add.s32 	%r174, %r47, 1;
	mul.lo.s32 	%r175, %r174, %r47;
	shr.u32 	%r176, %r175, 31;
	mad.lo.s32 	%r177, %r174, %r47, %r176;
	shr.s32 	%r178, %r177, 1;
	add.s32 	%r194, %r178, %r22;

BB0_21:
	mul.wide.s32 	%rd97, %r194, 4;
	add.s64 	%rd98, %rd32, %rd97;
	ld.global.f32 	%f74, [%rd98];
	.loc 2 71 1
	mul.f32 	%f75, %f74, %f46;
	sub.f32 	%f76, %f37, %f1;
	sub.f32 	%f77, %f38, %f2;
	sub.f32 	%f78, %f39, %f3;
	.loc 2 71 1
	fma.rn.f32 	%f79, %f75, %f76, %f34;
	fma.rn.f32 	%f80, %f75, %f77, %f35;
	fma.rn.f32 	%f81, %f75, %f78, %f36;

BB0_22:
	.loc 2 74 1
	st.global.f32 	[%rd1], %f79;
	.loc 2 75 1
	st.global.f32 	[%rd2], %f80;
	.loc 2 76 1
	st.global.f32 	[%rd3], %f81;
	.loc 2 23 1
	add.s32 	%r187, %r187, %r10;
	add.s32 	%r186, %r186, %r10;
	add.s32 	%r185, %r185, %r10;
	add.s32 	%r184, %r184, %r10;
	add.s32 	%r183, %r183, %r10;
	.loc 2 23 18
	add.s32 	%r188, %r188, 1;
	.loc 2 23 1
	setp.lt.s32 	%p13, %r188, %r57;
	@%p13 bra 	BB0_2;

BB0_23:
	.loc 2 78 2
	ret;
}


`
	addexchange_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry addexchange(
	.param .u64 addexchange_param_0,
	.param .u64 addexchange_param_1,
	.param .u64 addexchange_param_2,
	.param .u64 addexchange_param_3,
	.param .u64 addexchange_param_4,
	.param .u64 addexchange_param_5,
	.param .u64 addexchange_param_6,
	.param .u64 addexchange_param_7,
	.param .f32 addexchange_param_8,
	.param .f32 addexchange_param_9,
	.param .f32 addexchange_param_10,
	.param .u32 addexchange_param_11,
	.param .u32 addexchange_param_12,
	.param .u32 addexchange_param_13
)
{
	.reg .pred 	%p<14>;
	.reg .s16 	%rc<8>;
	.reg .s32 	%r<174>;
	.reg .f32 	%f<82>;
	.reg .s64 	%rd<99>;


	ld.param.u64 	%rd4, [addexchange_param_0];
	ld.param.u64 	%rd5, [addexchange_param_1];
	ld.param.u64 	%rd6, [addexchange_param_2];
	ld.param.u64 	%rd7, [addexchange_param_3];
	ld.param.u64 	%rd8, [addexchange_param_4];
	ld.param.u64 	%rd9, [addexchange_param_5];
	ld.param.u64 	%rd10, [addexchange_param_6];
	ld.param.u64 	%rd11, [addexchange_param_7];
	ld.param.f32 	%f46, [addexchange_param_8];
	ld.param.f32 	%f47, [addexchange_param_9];
	ld.param.f32 	%f48, [addexchange_param_10];
	ld.param.u32 	%r57, [addexchange_param_11];
	ld.param.u32 	%r58, [addexchange_param_12];
	ld.param.u32 	%r59, [addexchange_param_13];
	.loc 3 16 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 3 17 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 3 19 1
	setp.lt.s32 	%p1, %r8, %r59;
	setp.lt.s32 	%p2, %r4, %r58;
	and.pred  	%p3, %p2, %p1;
	.loc 3 23 1
	setp.gt.s32 	%p4, %r57, 0;
	.loc 3 19 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_23;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 36 1
	add.s32 	%r61, %r8, -1;
	mov.u32 	%r167, 0;
	.loc 4 238 5
	max.s32 	%r62, %r61, %r167;
	.loc 3 42 1
	add.s32 	%r63, %r59, -1;
	add.s32 	%r64, %r8, 1;
	.loc 4 210 5
	min.s32 	%r65, %r64, %r63;
	.loc 3 48 1
	add.s32 	%r66, %r4, -1;
	.loc 4 238 5
	max.s32 	%r67, %r66, %r167;
	.loc 3 54 1
	add.s32 	%r68, %r58, -1;
	add.s32 	%r69, %r4, 1;
	.loc 4 210 5
	min.s32 	%r70, %r69, %r68;
	.loc 3 23 1
	mad.lo.s32 	%r166, %r70, %r59, %r8;
	mul.lo.s32 	%r10, %r59, %r58;
	mad.lo.s32 	%r165, %r67, %r59, %r8;
	mad.lo.s32 	%r164, %r59, %r4, %r65;
	mad.lo.s32 	%r163, %r59, %r4, %r62;
	mad.lo.s32 	%r162, %r59, %r4, %r8;
	cvta.to.global.u64 	%rd22, %rd4;
	cvta.to.global.u64 	%rd23, %rd5;
	cvta.to.global.u64 	%rd24, %rd6;

BB2_2:
	cvta.to.global.u64 	%rd12, %rd11;
	cvta.to.global.u64 	%rd13, %rd7;
	.loc 3 27 1
	cvt.s64.s32 	%rd14, %r162;
	mul.wide.s32 	%rd15, %r162, 4;
	add.s64 	%rd16, %rd13, %rd15;
	ld.global.nc.f32 	%f1, [%rd16];
	cvta.to.global.u64 	%rd17, %rd8;
	.loc 3 27 1
	add.s64 	%rd18, %rd17, %rd15;
	ld.global.nc.f32 	%f2, [%rd18];
	cvta.to.global.u64 	%rd19, %rd9;
	.loc 3 27 1
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.nc.f32 	%f3, [%rd20];
	.loc 3 28 1
	add.s64 	%rd21, %rd12, %rd14;
	ld.global.u8 	%rc1, [%rd21];
	.loc 3 29 1
	add.s64 	%rd1, %rd22, %rd15;
	ld.global.f32 	%f4, [%rd1];
	add.s64 	%rd2, %rd23, %rd15;
	ld.global.f32 	%f5, [%rd2];
	add.s64 	%rd3, %rd24, %rd15;
	ld.global.f32 	%f6, [%rd3];
	.loc 3 37 1
	cvt.s64.s32 	%rd25, %r163;
	mul.wide.s32 	%rd26, %r163, 4;
	add.s64 	%rd27, %rd13, %rd26;
	ld.global.nc.f32 	%f7, [%rd27];
	add.s64 	%rd28, %rd17, %rd26;
	ld.global.nc.f32 	%f8, [%rd28];
	add.s64 	%rd29, %rd19, %rd26;
	ld.global.nc.f32 	%f9, [%rd29];
	.loc 3 38 1
	add.s64 	%rd30, %rd12, %rd25;
	ld.global.u8 	%rc2, [%rd30];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc2;
	setp.gt.s16 	%p6, %temp1, %temp2;
	}
	cvt.s32.s8 	%r21, %rc2;
	cvt.s32.s8 	%r22, %rc1;
	@%p6 bra 	BB2_4;

	add.s32 	%r78, %r22, 1;
	mul.lo.s32 	%r79, %r78, %r22;
	shr.u32 	%r80, %r79, 31;
	mad.lo.s32 	%r81, %r78, %r22, %r80;
	shr.s32 	%r82, %r81, 1;
	add.s32 	%r168, %r21, %r82;
	bra.uni 	BB2_5;

BB2_4:
	.loc 3 38 1
	add.s32 	%r83, %r21, 1;
	mul.lo.s32 	%r84, %r83, %r21;
	shr.u32 	%r85, %r84, 31;
	mad.lo.s32 	%r86, %r83, %r21, %r85;
	shr.s32 	%r87, %r86, 1;
	add.s32 	%r168, %r87, %r22;

BB2_5:
	cvta.to.global.u64 	%rd32, %rd10;
	.loc 3 38 1
	mul.wide.s32 	%rd33, %r168, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.f32 	%f49, [%rd34];
	.loc 3 39 1
	mul.f32 	%f50, %f49, %f48;
	sub.f32 	%f51, %f7, %f1;
	sub.f32 	%f52, %f8, %f2;
	sub.f32 	%f53, %f9, %f3;
	.loc 3 39 1
	fma.rn.f32 	%f10, %f50, %f51, %f4;
	fma.rn.f32 	%f11, %f50, %f52, %f5;
	fma.rn.f32 	%f12, %f50, %f53, %f6;
	.loc 3 43 1
	cvt.s64.s32 	%rd36, %r164;
	mul.wide.s32 	%rd37, %r164, 4;
	add.s64 	%rd38, %rd13, %rd37;
	ld.global.nc.f32 	%f13, [%rd38];
	add.s64 	%rd40, %rd17, %rd37;
	ld.global.nc.f32 	%f14, [%rd40];
	add.s64 	%rd42, %rd19, %rd37;
	ld.global.nc.f32 	%f15, [%rd42];
	.loc 3 44 1
	add.s64 	%rd43, %rd12, %rd36;
	ld.global.u8 	%rc3, [%rd43];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc3;
	setp.gt.s16 	%p7, %temp1, %temp2;
	}
	cvt.s32.s8 	%r27, %rc3;
	@%p7 bra 	BB2_7;

	add.s32 	%r90, %r22, 1;
	mul.lo.s32 	%r91, %r90, %r22;
	shr.u32 	%r92, %r91, 31;
	mad.lo.s32 	%r93, %r90, %r22, %r92;
	shr.s32 	%r94, %r93, 1;
	add.s32 	%r169, %r27, %r94;
	bra.uni 	BB2_8;

BB2_7:
	.loc 3 44 1
	add.s32 	%r95, %r27, 1;
	mul.lo.s32 	%r96, %r95, %r27;
	shr.u32 	%r97, %r96, 31;
	mad.lo.s32 	%r98, %r95, %r27, %r97;
	shr.s32 	%r99, %r98, 1;
	add.s32 	%r169, %r99, %r22;

BB2_8:
	mul.wide.s32 	%rd46, %r169, 4;
	add.s64 	%rd47, %rd32, %rd46;
	ld.global.f32 	%f54, [%rd47];
	.loc 3 45 1
	mul.f32 	%f55, %f54, %f48;
	sub.f32 	%f56, %f13, %f1;
	sub.f32 	%f57, %f14, %f2;
	sub.f32 	%f58, %f15, %f3;
	.loc 3 45 1
	fma.rn.f32 	%f16, %f55, %f56, %f10;
	fma.rn.f32 	%f17, %f55, %f57, %f11;
	fma.rn.f32 	%f18, %f55, %f58, %f12;
	.loc 3 49 1
	cvt.s64.s32 	%rd49, %r165;
	mul.wide.s32 	%rd50, %r165, 4;
	add.s64 	%rd51, %rd13, %rd50;
	ld.global.nc.f32 	%f19, [%rd51];
	add.s64 	%rd53, %rd17, %rd50;
	ld.global.nc.f32 	%f20, [%rd53];
	add.s64 	%rd55, %rd19, %rd50;
	ld.global.nc.f32 	%f21, [%rd55];
	.loc 3 50 1
	add.s64 	%rd56, %rd12, %rd49;
	ld.global.u8 	%rc4, [%rd56];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc4;
	setp.gt.s16 	%p8, %temp1, %temp2;
	}
	cvt.s32.s8 	%r32, %rc4;
	@%p8 bra 	BB2_10;

	add.s32 	%r102, %r22, 1;
	mul.lo.s32 	%r103, %r102, %r22;
	shr.u32 	%r104, %r103, 31;
	mad.lo.s32 	%r105, %r102, %r22, %r104;
	shr.s32 	%r106, %r105, 1;
	add.s32 	%r170, %r32, %r106;
	bra.uni 	BB2_11;

BB2_10:
	.loc 3 50 1
	add.s32 	%r107, %r32, 1;
	mul.lo.s32 	%r108, %r107, %r32;
	shr.u32 	%r109, %r108, 31;
	mad.lo.s32 	%r110, %r107, %r32, %r109;
	shr.s32 	%r111, %r110, 1;
	add.s32 	%r170, %r111, %r22;

BB2_11:
	mul.wide.s32 	%rd59, %r170, 4;
	add.s64 	%rd60, %rd32, %rd59;
	ld.global.f32 	%f59, [%rd60];
	.loc 3 51 1
	mul.f32 	%f60, %f59, %f47;
	sub.f32 	%f61, %f19, %f1;
	sub.f32 	%f62, %f20, %f2;
	sub.f32 	%f63, %f21, %f3;
	.loc 3 51 1
	fma.rn.f32 	%f22, %f60, %f61, %f16;
	fma.rn.f32 	%f23, %f60, %f62, %f17;
	fma.rn.f32 	%f24, %f60, %f63, %f18;
	.loc 3 55 1
	cvt.s64.s32 	%rd62, %r166;
	mul.wide.s32 	%rd63, %r166, 4;
	add.s64 	%rd64, %rd13, %rd63;
	ld.global.nc.f32 	%f25, [%rd64];
	add.s64 	%rd66, %rd17, %rd63;
	ld.global.nc.f32 	%f26, [%rd66];
	add.s64 	%rd68, %rd19, %rd63;
	ld.global.nc.f32 	%f27, [%rd68];
	.loc 3 56 1
	add.s64 	%rd69, %rd12, %rd62;
	ld.global.u8 	%rc5, [%rd69];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc5;
	setp.gt.s16 	%p9, %temp1, %temp2;
	}
	cvt.s32.s8 	%r37, %rc5;
	@%p9 bra 	BB2_13;

	add.s32 	%r114, %r22, 1;
	mul.lo.s32 	%r115, %r114, %r22;
	shr.u32 	%r116, %r115, 31;
	mad.lo.s32 	%r117, %r114, %r22, %r116;
	shr.s32 	%r118, %r117, 1;
	add.s32 	%r171, %r37, %r118;
	bra.uni 	BB2_14;

BB2_13:
	.loc 3 56 1
	add.s32 	%r119, %r37, 1;
	mul.lo.s32 	%r120, %r119, %r37;
	shr.u32 	%r121, %r120, 31;
	mad.lo.s32 	%r122, %r119, %r37, %r121;
	shr.s32 	%r123, %r122, 1;
	add.s32 	%r171, %r123, %r22;

BB2_14:
	mul.wide.s32 	%rd71, %r171, 4;
	add.s64 	%rd72, %rd32, %rd71;
	ld.global.f32 	%f64, [%rd72];
	.loc 3 57 1
	mul.f32 	%f65, %f64, %f47;
	sub.f32 	%f66, %f25, %f1;
	sub.f32 	%f67, %f26, %f2;
	sub.f32 	%f68, %f27, %f3;
	.loc 3 57 1
	fma.rn.f32 	%f79, %f65, %f66, %f22;
	fma.rn.f32 	%f80, %f65, %f67, %f23;
	fma.rn.f32 	%f81, %f65, %f68, %f24;
	setp.eq.s32 	%p10, %r57, 1;
	.loc 3 60 1
	@%p10 bra 	BB2_22;

	.loc 3 62 1
	add.s32 	%r125, %r167, -1;
	mov.u32 	%r126, 0;
	.loc 4 238 5
	max.s32 	%r127, %r125, %r126;
	.loc 3 62 1
	mad.lo.s32 	%r128, %r127, %r58, %r4;
	mad.lo.s32 	%r129, %r128, %r59, %r8;
	.loc 3 63 1
	cvt.s64.s32 	%rd74, %r129;
	mul.wide.s32 	%rd76, %r129, 4;
	add.s64 	%rd77, %rd13, %rd76;
	ld.global.nc.f32 	%f31, [%rd77];
	add.s64 	%rd79, %rd17, %rd76;
	ld.global.nc.f32 	%f32, [%rd79];
	add.s64 	%rd81, %rd19, %rd76;
	ld.global.nc.f32 	%f33, [%rd81];
	.loc 3 64 1
	add.s64 	%rd82, %rd12, %rd74;
	ld.global.u8 	%rc6, [%rd82];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc6;
	setp.gt.s16 	%p11, %temp1, %temp2;
	}
	cvt.s32.s8 	%r42, %rc6;
	@%p11 bra 	BB2_17;

	add.s32 	%r131, %r22, 1;
	mul.lo.s32 	%r132, %r131, %r22;
	shr.u32 	%r133, %r132, 31;
	mad.lo.s32 	%r134, %r131, %r22, %r133;
	shr.s32 	%r135, %r134, 1;
	add.s32 	%r172, %r42, %r135;
	bra.uni 	BB2_18;

BB2_17:
	.loc 3 64 1
	add.s32 	%r136, %r42, 1;
	mul.lo.s32 	%r137, %r136, %r42;
	shr.u32 	%r138, %r137, 31;
	mad.lo.s32 	%r139, %r136, %r42, %r138;
	shr.s32 	%r140, %r139, 1;
	add.s32 	%r172, %r140, %r22;

BB2_18:
	mul.wide.s32 	%rd85, %r172, 4;
	add.s64 	%rd86, %rd32, %rd85;
	ld.global.f32 	%f69, [%rd86];
	.loc 3 65 1
	mul.f32 	%f70, %f69, %f46;
	sub.f32 	%f71, %f31, %f1;
	sub.f32 	%f72, %f32, %f2;
	sub.f32 	%f73, %f33, %f3;
	.loc 3 65 1
	fma.rn.f32 	%f34, %f70, %f71, %f79;
	fma.rn.f32 	%f35, %f70, %f72, %f80;
	fma.rn.f32 	%f36, %f70, %f73, %f81;
	.loc 3 68 1
	add.s32 	%r142, %r57, -1;
	add.s32 	%r143, %r167, 1;
	.loc 4 210 5
	min.s32 	%r144, %r143, %r142;
	.loc 3 68 1
	mad.lo.s32 	%r145, %r144, %r58, %r4;
	mad.lo.s32 	%r146, %r145, %r59, %r8;
	.loc 3 69 1
	cvt.s64.s32 	%rd87, %r146;
	mul.wide.s32 	%rd89, %r146, 4;
	add.s64 	%rd90, %rd13, %rd89;
	ld.global.nc.f32 	%f37, [%rd90];
	add.s64 	%rd92, %rd17, %rd89;
	ld.global.nc.f32 	%f38, [%rd92];
	add.s64 	%rd94, %rd19, %rd89;
	ld.global.nc.f32 	%f39, [%rd94];
	.loc 3 70 1
	add.s64 	%rd95, %rd12, %rd87;
	ld.global.u8 	%rc7, [%rd95];
	{
	.reg .s16 	%temp1;
	.reg .s16 	%temp2;
	cvt.s16.s8 	%temp1, %rc1;
	cvt.s16.s8 	%temp2, %rc7;
	setp.gt.s16 	%p12, %temp1, %temp2;
	}
	cvt.s32.s8 	%r47, %rc7;
	@%p12 bra 	BB2_20;

	add.s32 	%r148, %r22, 1;
	mul.lo.s32 	%r149, %r148, %r22;
	shr.u32 	%r150, %r149, 31;
	mad.lo.s32 	%r151, %r148, %r22, %r150;
	shr.s32 	%r152, %r151, 1;
	add.s32 	%r173, %r47, %r152;
	bra.uni 	BB2_21;

BB2_20:
	.loc 3 70 1
	add.s32 	%r153, %r47, 1;
	mul.lo.s32 	%r154, %r153, %r47;
	shr.u32 	%r155, %r154, 31;
	mad.lo.s32 	%r156, %r153, %r47, %r155;
	shr.s32 	%r157, %r156, 1;
	add.s32 	%r173, %r157, %r22;

BB2_21:
	mul.wide.s32 	%rd97, %r173, 4;
	add.s64 	%rd98, %rd32, %rd97;
	ld.global.f32 	%f74, [%rd98];
	.loc 3 71 1
	mul.f32 	%f75, %f74, %f46;
	sub.f32 	%f76, %f37, %f1;
	sub.f32 	%f77, %f38, %f2;
	sub.f32 	%f78, %f39, %f3;
	.loc 3 71 1
	fma.rn.f32 	%f79, %f75, %f76, %f34;
	fma.rn.f32 	%f80, %f75, %f77, %f35;
	fma.rn.f32 	%f81, %f75, %f78, %f36;

BB2_22:
	.loc 3 74 1
	st.global.f32 	[%rd1], %f79;
	.loc 3 75 1
	st.global.f32 	[%rd2], %f80;
	.loc 3 76 1
	st.global.f32 	[%rd3], %f81;
	.loc 3 23 1
	add.s32 	%r166, %r166, %r10;
	add.s32 	%r165, %r165, %r10;
	add.s32 	%r164, %r164, %r10;
	add.s32 	%r163, %r163, %r10;
	add.s32 	%r162, %r162, %r10;
	.loc 3 23 18
	add.s32 	%r167, %r167, 1;
	.loc 3 23 1
	setp.lt.s32 	%p13, %r167, %r57;
	@%p13 bra 	BB2_2;

BB2_23:
	.loc 3 78 2
	ret;
}


`
)
