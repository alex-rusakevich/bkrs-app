#include <windows.h>
#include <winuser.h>
#include <uxtheme.h>
#include <stdio.h>

#define MOD_NOREPEAT 0x4000

static inline void moveAppWindow(int x, int y) {
	SetWindowPos(GetActiveWindow(), 0, x, y, 0, 0, SWP_NOSIZE);
}

static inline void makeWindowAlwaysOnTop() {
	HWND hwnd = GetActiveWindow();
	SetWindowPos(hwnd, HWND_TOPMOST, 0, 0, 0, 0, SWP_NOSIZE | SWP_NOMOVE);
}

static inline void loadIcon() {
	HICON smallIcon, bigIcon;

    smallIcon = (HICON)LoadImageA(0, "icon.ico", IMAGE_ICON, 16, 16, LR_LOADFROMFILE);
    bigIcon = (HICON)LoadImageA(0, "icon.ico", IMAGE_ICON, 32, 32, LR_LOADFROMFILE);

    SendMessage(GetActiveWindow(), WM_SETICON, ICON_BIG, (LPARAM)bigIcon);
	SendMessage(GetActiveWindow(), WM_SETICON, ICON_SMALL, (LPARAM)smallIcon);
}

static inline void minimizeOnShortcut(HWND hwnd) {
    RegisterHotKey(NULL, 1, MOD_ALT | MOD_NOREPEAT, 0x53); // Alt+S

    MSG msg = {0};
    while (GetMessage(&msg, NULL, 0, 0) != 0) {
        if (msg.message == WM_HOTKEY) {
            if((int)msg.wParam == 1) {
                if(IsIconic(hwnd)) {
                    ShowWindow(hwnd, SW_RESTORE);
                } else {
                    ShowWindow(hwnd, SW_MINIMIZE);
                }
            }
        }
    }
}
