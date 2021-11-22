;----------------------------------------------------------------------------
;				NES Startup
;					Copyright (C) 2007, Tekepen
;----------------------------------------------------------------------------
.setcpu		"6502"
.autoimport	on
.IMPORTZP	sp

.global	_NesMain

; iNES�w�b�_
.segment "HEADER"
	.byte	$4E, $45, $53, $1A	; "NES" Header
	.byte	$02			; PRG-BANKS
	.byte	$01			; CHR-BANKS
	.byte	$01			; Vetrical Mirror
	.byte	$00			; 
	.byte	$00, $00, $00, $00	; 
	.byte	$00, $00, $00, $00	; 

.segment "STARTUP"
; ���Z�b�g���荞��
.proc	Reset
	sei
	ldx	#$ff
	txs
	
	; �\�t�g�E�F�A�X�^�b�N�ݒ�
	lda	#$ff
	sta	sp
	lda	#$03
	sta	sp + 1
	
	jsr	_NesMain

.endproc

.segment "VECINFO"
	.word	$0000
	.word	Reset
	.word	$0000

; �p�^�[���e�[�u��
.segment "CHARS"
	.incbin	"character.chr"
