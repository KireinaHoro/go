TEXT	main(SB),512|7,$0
	CMP	RT1, R9
	SUBCC	RT1, R9, ZR
	BLE	ICC, label
	MOVD	$1, RT1
	RET
label:
	MOVD	$2, RT1
	CMP	$0, R10
	CMP ZR, R11
	CMP	$42, R8
	SUBCC	$42, R12, ZR
	RET