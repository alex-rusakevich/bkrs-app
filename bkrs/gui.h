#define WIN32_LEAN_AND_MEAN

#include <windows.h>
#include <winuser.h>
#include <stdio.h>

#define MOD_NOREPEAT 0x4000
#define TOGGLE_MINIMIZE_KEYS MOD_ALT | MOD_NOREPEAT, 0x44 /* Alt+D */

HWND hwnd;

static inline void initGUI(HWND hwndIn) {
    hwnd = hwndIn;
}

static inline void moveAppWindow(int x, int y) {
	SetWindowPos(hwnd, 0, x, y, 0, 0, SWP_NOSIZE);
}

static inline void makeWindowAlwaysOnTop() {
	SetWindowPos(hwnd, HWND_TOPMOST, 0, 0, 0, 0, SWP_NOSIZE | SWP_NOMOVE);
}

static inline void loadIcon() {
	HICON smallIcon, bigIcon;

    smallIcon = (HICON)LoadImageA(0, "icon.ico", IMAGE_ICON, 16, 16, LR_LOADFROMFILE);
    bigIcon = (HICON)LoadImageA(0, "icon.ico", IMAGE_ICON, 32, 32, LR_LOADFROMFILE);

    SendMessage(hwnd, WM_SETICON, ICON_BIG, (LPARAM)bigIcon);
	SendMessage(hwnd, WM_SETICON, ICON_SMALL, (LPARAM)smallIcon);
}

static inline void minimizeOnShortcut() {
    RegisterHotKey(NULL, 1, TOGGLE_MINIMIZE_KEYS); // Alt+A

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
