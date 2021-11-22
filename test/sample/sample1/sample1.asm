;----------------------------------------------------------------------------
;				NES Sample1 "HELLO, WORLD!"
;					Copyright (C) 2007, Tekepen
;----------------------------------------------------------------------------
.setcpu		"6502"
.autoimport	on

; iNES�w�b�_
.segment "HEADER"
	.byte	$4E, $45, $53, $1A	; "NES" Header
	.byte	$02			; PRG-BANKS
	.byte	$01			; CHR-BANKS
	.byte	$01			; Vetrical Mirror
	.byte	$00			;
	.byte	$00, $00, $00, $00	;a
	.byte	$00, $00, $00, $00	;

.segment "STARTUP"
; ���Z�b�g���荞��
.proc	Reset
	sei
	ldx	#$ff
	txs

; �X�N���[���I�t
	lda	#$00
	sta	$2000
	sta	$2001

; �p���b�g�e�[�u���֓]��(BG�p�̂ݓ]��)
	lda	#$3f
	sta	$2006
	lda	#$00
	sta	$2006
	ldx	#$00
	ldy	#$10
copypal:
	lda	palettes, x
	sta	$2007
	inx
	dey
	bne	copypal

; �l�[���e�[�u���֓]��(��ʂ̒����t��)
	lda	#$21
	sta	$2006
	lda	#$c9
	sta	$2006
	ldx	#$00
	ldy	#$0d		; 13�����\��
copymap:
	lda	string, x
	sta	$2007
	inx
	dey
	bne	copymap

; �X�N���[���ݒ�
	lda	#$00
	sta	$2005
	sta	$2005

; �X�N���[���I��
	lda	#$08
	sta	$2000
	lda	#$1e
	sta	$2001

; �������[�v
mainloop:
	jmp	mainloop
.endproc

; �p���b�g�e�[�u��
palettes:
	.byte	$0f, $00, $10, $20
	.byte	$0f, $06, $16, $26
	.byte	$0f, $08, $18, $28
	.byte	$0f, $0a, $1a, $2a

; �\��������
string:
	.byte	"HELLO, WORLD!"

.segment "VECINFO"
	.word	$0000
	.word	Reset
	.word	$0000

; �p�^�[���e�[�u��
.segment "CHARS"
	.incbin	"character.chr"
