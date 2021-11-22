/*---------------------------------------------------------------------------
				NES Sample2 "HELLO, WORLD!" - C Version
											Copyright (C) 2007, Tekepen
---------------------------------------------------------------------------*/

// �X�N���[���̃I���E�I�t
void ShowScreen(char ShowFlag)
{
	if (ShowFlag == 1) {
		*(char*)0x2000 = 0x08;
		*(char*)0x2001 = 0x1e;
	} else {
		*(char*)0x2000 = 0x00;
		*(char*)0x2001 = 0x00;
	}
}

// �X�N���[���ݒ�
void SetScroll(unsigned char x, unsigned char y)
{
	*(char*)0x2005 = x;
	*(char*)0x2005 = y;
}

// ���C���֐�
char NesMain()
{
	const char palettes[] = {
		0x0f, 0x00, 0x10, 0x20,
		0x0f, 0x06, 0x16, 0x26,
		0x0f, 0x08, 0x18, 0x28,
		0x0f, 0x0a, 0x1a, 0x2a
	};
	const char string[] = "HELLO, WORLD!";
	char i;

	ShowScreen(0);

	// �p���b�g����������
	*(char*)0x2006 = 0x3f;
	*(char*)0x2006 = 0x00;
	for (i = 0; i < 0x10; i ++)
		*(char*)0x2007 = palettes[i];

	// �l�[���e�[�u���֏�������
	*(char*)0x2006 = 0x21;
	*(char*)0x2006 = 0xc9;
	for (i = 0; i < 13; i ++)
		*(char*)0x2007 = string[i];

	SetScroll(0, 0);
	ShowScreen(1);

	// �������[�v
	while (1);

	return 0;
}
