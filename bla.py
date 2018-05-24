import pyperclip
import time
import os
while True:
    a = pyperclip.paste()
    time.sleep(0.3)
    if a != pyperclip.paste():
        if "http" in pyperclip.paste():
            if "youtube" in pyperclip.paste():
                os.system("mpv " +  pyperclip.paste())
            else:
                os.system("google-chrome " + pyperclip.paste())
