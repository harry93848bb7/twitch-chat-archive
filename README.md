# Twitch Chat Archiver

A tool to archive an entire VOD chat. The goal of this tool is to save all relavent data for a complete offline replay of a VOD chat.

### Features
- Archives include embedded badges and emote images / gifs
- Supported JSON and Protocol Buffer output files.

### Supported Image / GIF Archives
- Twitch Robot Emotes
- Twitch Global Emotes
- Channel Subscription Emotes
- BetterTTV Global Emotes
- BetterTTV Channel Emotes
- FrankerFaceZ Global Emotes
- FrankerFaceZ Channel Emotes
- Twitch Global Badges
- Channel Subscription Badges
- Channel Bit Badges
- Channel Cheer Badges

## Installation / Usage
You will need a Twitch Client ID. Register an application [here](https://dev.twitch.tv/console/apps/create) and copy the Client ID. The OAuth Redirect URL can be any value.


1. Download the [latest release](https://github.com/harry93848bb7/twitch-chat-archive/releases) for your system
2. Extract the files in a folder of your choice
2. Run the CLI: `./twitch-chat-archiver -client_id=4ab086186863bffbf81a73d359fc2ec7 -vod_id=951894217`
<details><summary>Archive Output Sample</summary>
<p>

```go
{
   "vod_id": 951894217,
   "title": "IM BACK--WE GOT A LOT TO TALK ABOUT\n",
   "category": "Just Chatting",
   "length": {
      "seconds": 9629
   },
   "recorded_at": {
      "seconds": 1615927383
   },
   "channel": {
      "name": "richwcampbell",
      "id": 127463427
   },
   "badges": [
      {
         "code": "premium",
         "version": "1",
         "title": "Prime Gaming",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABIAAAASCAYAAABWzo5XAAAA+0lEQVR4nGJhmHbTk4Hx3yyG//9lGMgBjIxPGP4zpTEyTL/+mGxDkAxjotgQEPj/X4aJYkOggDYGMWJRUGMszFBpKIwhjq4WxaAptuIoksX6QgzNZmIMbRZiDAV6QihyU20lcBuUpSPE0GslBmFrCzL0WCEM7rcWZ0jXEoCwrcQYMrQFUAxiQeb8//+foUhfmEFDgJ3BU44bwzvT7SQY/BR4GbzkecBqcRoEAyCF2AAjIyNcDsTG6bX/WLUTB1AM+v2PeKP+oqlFMWjJrU8YCrCBf///M6y6+wlFjJFh2jVKfITdRYPEIFB5QimAFCNMaRQZBi3YAAEAAP//NaBCmbt+SO0AAAAASUVORK5CYII="
      },
      {
         "code": "sub-gifter",
         "version": "10",
         "title": "10 Gift Subs",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABIAAAASCAYAAABWzo5XAAAB0UlEQVR4nIyRv2sUQRTHv2/2gnamSGOjEQRBEE6rFEHOdFYiiK13ZUA4/AtOOysPG9NETq0tksrOWARJ5Qm2wu1VKgSMkOju7Mx7Mvsru5fZXB5X7Mz73vd95vtaqNTa6PcACl0CDoTty4+9pTf1/n6XKOiDqE2Qr4aj+596F0PXU4Xo9uaPR9bGXTbJM6vNW7Z20Nn8udMZTRbT/utfA2v4hbV623JyxyQazNQv/l8SidZtJgp31y+lFJ2NyZYV2mFSw9WNaShx1NXEt/bWr6QEq69CEKikLb9WhpNlUuYLCW+DZcsIhUrxH6WCsYgsQqmOHB1OaeFcGy3VB6nLwgtre08yY6pmsDIcL7MNngK4B9AFghAE7pePJfd5EADv6Hww/Pz4+rRs3X0vAk8xW+x/H5fYTrR09SaUCnxytLy3bgsqgE2iYyNBo8mpRq6sjgrw02RnMYrP4pEZfXhANPrmz+n5blQ7P7zmN+ndIEqJDpMGoiSunWd1xRJQPO3IY0S1jLILnw5Vo78NAqMjEB2H5NPNJXLFRkNE4LxcivOJTDNRdXSTrjT61xT2TEZOV13vibCLSdUGPFubT2TqU/KX5GHnB8p0s8MK/f8AAAD//1s03EsOQsx/AAAAAElFTkSuQmCC"
      },
      {
         "code": "sub-gifter",
         "version": "25",
         "title": "25 Gift Subs",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABIAAAASCAYAAABWzo5XAAAB2ElEQVR4nJRTu24TQRQ917O4gALTGYfHghJDg2TxBcAPIL4AVxQ0LmjoDB0SggjxA8AXJF+QPwAkCoRA8gYUBZGCxI+ItXfmoIl3dmfDLIjb7O7cc889c+7dCF4cdDtDQvog9gm8OPN555Wf/7l2vi/gQIAegPeHi+zO2eR7YnMNB/rRXbmbEv2MeKwbeK3B4d7aytYobrVsfq/bGWqY5wvBpiFuzkmoSA1cfaFobkwPaCTnvixV7F5tb+iFbJ1Qp9Z3Vk8mqWZ/oc31S7mCb6sdCFioFfcyitsxlbw1kE1qvWGESZONA0byzhAtJbyBDNtaSY+CgQIvSoZbjlh8Dz7G7RgKjwi5LeBpiAhZ6Wo/9gXypmmwfnm0u13keOUCEQhN4kM2L3UTuBY1oURCcETBU+CoYA7CKhLrBlFL8lciG6m7FoOi/4PITeXfPIjk01fhsydBaPryafXg3v0giTx4KEtFk3Gw6x+KxmOUB1W/oiqgGnZ7nT1ijZ6EcSXRZBJM/jJEMSjW40qiaRhgx7+8hV1FU4vziKbBZOGRyZ8VopBHlkg8D/O69Pj+zGb5TknJlWOWRIdTr0s5KUdkjT7652ae8mNNo6JTyCNb7HWtw9n4HQAA//8I2N1ZmK9/QgAAAABJRU5ErkJggg=="
      },
      {
         "code": "turbo",
         "version": "1",
         "title": "Turbo",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABIAAAASCAYAAABWzo5XAAAA9klEQVR4nGKJtJ4Tz/jv34T/DAwCDOSBBwxMTA1MDP/+NlBgCAgoMPz7N4GJgYFRgQJDYECAiQqGgAELjLHsWCpJGqOsZqPwCbro+vnnRBmM16CHt98yPLj1hjKDXj//zNBXsZvBxA57XICCAjk4WHAZBDIEZFh+8Aow3yNMhyGuwJI0Fy2eeBzsLRiQVxHCawhOg2LzLcHObl8QxCAiwcNQ1OmG1xAGfF4DgUPbbzEUd7oxiEryYsiRFP2aBpIM8qrCBF1D0CATe+JzD9xr6E4lFVAtrzExMjB8oII5D5j+/f9byMDw/wG5JoAc8v8/QyMgAAD//0Y6RlkIka6FAAAAAElFTkSuQmCC"
      },
      {
         "code": "moderator",
         "version": "1",
         "title": "Moderator",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABIAAAASCAYAAABWzo5XAAABNklEQVR4nGJhWMccz/CfYQIDA4MAA3ngAQMjQwMjw1rm+wwMDApkGgIDH1iINSRfJY9hgl4fhjjjOpARDAJMxBiiwCXP0KBZh1cNQYMEWAUY9tvuBdMUGVSvWcegwE3Y93gNAoVLgUoeQUPwGkQoXD78/sBQeKkIzmdBd8HGZxsZPvz+iDdcHnx7yOB4yAlMI8Ba5v8gXHCx8D8I3P/64P/8Bwv+4wITbk/8L7BJ+D9MHwwzghlQr4BcgStgQV5pvN7EMOHOJKzycIPwGYbdK6iAYPQveLiIwXCvMV5DUAxCdg1I04OvD8DiDqL2DAKs/ITsQxg033ge3BCQNxwPO4MNA1kQIOlP0CCUwAYZlng2Ce4NsCFS/jgDGN2g9xSURVDw/wHIa4UgBgWGfGBgYG4EBAAA//+kk5ttosPD1AAAAABJRU5ErkJggg=="
      }
   ],
   "emotes": [
      {
         "code": ":tf:",
         "source": "BetterTTV Global Emotes",
         "image_type": "png",
         "image_data": "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"
      },
      {
         "code": "D:",
         "source": "BetterTTV Global Emotes",
         "image_type": "png",
         "image_data": "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"
      },
      {
         "code": "KKona",
         "source": "BetterTTV Global Emotes",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABkAAAAiCAYAAACqVHINAAAJrklEQVR4nHyWaaxc91nGf2efs82ZuXPnbrZru2lt0sRNihUbNxCVJEAVlQZFlKAitUTQSlUQfEgVgUSluCzqh1ApAkr4ALQglkKbklZNaIJpnTapGzdO4ni5tnNz79xt7jLrmbOvaMZCUFL3/TjzP89z3vf/Pud5ZG5Q5577N4TaAkkQUIolkmJ81DDVjwz6/SNZmhtaRXkNWX8yDOOvG4bO5bPf5eMPP4ogCG/DevsvQNbtsHz2AvGeGsNh79Nlnn3aNJ3ZwOuysb5GWYKqKiiqRoGwpWnq14ZD//NebLx58+G9nDhx4kfwpP9PkLdfo/XGN+ka9eNFkX8v9N1fCzzPEijYWmtRpBm+OyIOPbI0ZrvdtsL+4I40jX83CodMTx84rUoxPzz3OidPnpxg/si4yrIkA/zN6MPD3c2ni7xPGkYosk5nZ5sgiJlqzqDqNq7bxR+5ZFHIpruNVlERFPmx1eVXb//snzz+wJNPfK4c443HJ/xfgo8JAp9Zunqi0155qb+zO/kzjiOqtRqeF1Cvz2HXHCRJZDjos9PepLO7Tq/bmZwrx28tKbzz8C1f/OjHPvmQ6yU4tnadZEywunqNNBfVzvb6stvdXkjTcnJjum5jmhaqrmM7VXy/z+rSZbrdAY7ToFqtsdJaYmN5hTSNURQFw7azhb0HnSCeDX7rE/dfJwnDkEKocOHVFx6PwtEjkR8gKSqGaSFrOoZuEAZDLp8/y4svnWb5yhKBH2LVLG4/eoz33naUOE7Z3tokiUIUUcM0zbs73fDbMzO3IERRztZWi6wQnN2t1nYw6mvjvlXNQDMtRDJ22i2e/dpXuHbpAnKaU5FUxosQ5CVelmI1m9xx/Dizc3tIowjf8xFF8cm7jvGppdUCcbff5Rtffpw49D+kKrImCiJFkU/2TiwTlhdf4R//9gu8fPo7lEMXHXAsnUajxpSlMy0JuK0VTj3zLGdeOk1Wpui2jaTm93cGPmWZI0a+z6ce/UvSLP0AZUFRpNc3IQ1445VT/NOX/o7FVy6hlQpGxURWFOIkwXVHhEGAKEpYWgV/t8frZ1/nwhsXYKyeQppvb6m/9NwPRshFGU0uxh1sHOjtrBMFo/EBWmtLnHrueRYX16jKCrNNm+ZMk7rpIJYFQRwgSBKOaZFlGRvdDpdW1njzzWvUHBvTskmSyq/8wvtPfEtWFAcROHPmhW+8denivWWc0u91WN/YYnVrSCEoJHKFtWFCWPb4qX06+xsOTcFGkEV6fkCr12er55IiEe702G1vMdJ6yIrY+v0/uAd5a/UsB/d/mAWr8R/KTUd3p+tmc3P5CtNOi0PvCHA0jUpFI8tS4iREEDLSPMEydEqhJIg9EEpmpqbYtzBPmhWYgkKWlLlXeH9+9wf/GOH7T/4OkhLBbv1KeuDwIWfvDFoRce3iea4tXeXca5e5tDIgAmYtlapSsH9+FlOWyCVYbu/SHgTodpNb90+zf36aUtLwUpEFXTp5123JY9Inb5micDfN9TNrf2a+532Ihkjqu8S+j6ppVO06FBkSJVk4IkwCNFVGQqQ3Cmh3B4zZKtU6UZQQRAPqNZuaNUVx5fydldrgc7Igz5DKqr3vzltRLZE4DjFkGcOqolc0ZiwdLRiyoknsDFRG4ZCqbnLkwLtY2VwlTjIUVWahaWBUdIaRR+LHNOySYOTHZeKksvrud6JWg6JcL1l6/ln2fPA+FEPBsacoRAE5Sjl0U8lsc4AfxwRJydGfvQdLyBD+8xkKRcWLU0ZhSBzEOLUGU1UHhIwtZybeWmog6nsr2LNKgpBj7X8Pgqbg+wG5UDAzcwCrOT9R/nB7k9biRQ4dvoXbjx3Dqk2hWw5V22G+NkVDt9AUnfmFeQzDJI1TelnaW+lHyCJVsjSQpw+9F101GUR9otgnESW6vXVGmxt0VlsEbkR3lJLEHj987quMej3C0Qi1zAnSmKokI+sKQhAiGDKKriBb6mVRnUKuGDalmLsVoxH7aaqVfobvjyAr6a+usb3Wwt118cKM5sw8L754hvWdNg29gp8kiLKEY1TJy5iaqiCpEoIgURYxtl1rHXng98YjX8auDZL1q9mjZXXfE2U51r9IELhsrrZYWVplN4hw6g0cSaCQNWqzc6xud5HlysSKDblClscMBn32SBKGkBAWEbpde+uhe48gu8ImpfsQ1t7N77iRCHFOUYpEWUEuKwzjhI4/opQqlFlOw1BRigJDFLEqKpIkoakiXiKx3nWJ0pgDlQW6Q49eGrx832EbMeYXkaWQQrL3JEFEGMaUIpQySJpEc2EWGZFzr15kdbeHm2WEpcyUWcNSNWpmheXNbc5fuIoXemOXIy1FhlHSu3hh8aW8uoCsGTOU8Q6CaLfT3GPkD8iTiDyOiOJ8Yrf7987TdBpEGUSyhDm3QBqHSLpEv9entd6mUTWYm2uwMDtNGCRIiv5UO2tySNeRzy/GHNBd/DS6KmX9ZOQO1CJJ8b0YL7remaKqHFyoYssWtqKji1Xe6u6wsdMhTRJ+5tBBqk2biHwizr47ZJjy7EH3ZU786h8iPf3U37Pwjjv5r3PDtC6t/3yZeAc9z6Pvuuz2XQZugCaJVISxW8pYFYXu7uZEmGkpYFcrOI6OrKkM/ISOO8ILw6he7z+s1pvxsduuBxK0qZuJuo/woXv/+v47Du/7d8QcP0gIgoDeKMZUJRpWFVPIKPKMq1t9Oq5PRVbY17Rp1B38XKI9dAnDiCKJvnLiuP+R0DM5+VcX/zcSNZozdHd3uO+uY1+/9d1zv5wlGaPQJctKBl7C7NwCtx2cYU9zD2Gasb6xRr/rUooJnVHEWzsDNAVIYsxKdO+Xn/dOffepm/m5B7498atJjQlOPvIgz7xw7sFzl6+dHoYBBSWSWKDKJRvbHl5pMkJip9NhbbtPMlaUYhNmCkmcIBUFipQ8f/xo9dRnH943IfixWXicwaZqqvTTR276m5quf1ymwLBMum6MWZvF1HSiOCAMYoRx4FAEkiyedCDkITPTxeF/+OqVq/32H6HPf+bGgfuu972Lyxvr7G00PlHX1c+btmoZeoU0Kxn6GUGYkJfCJMg1HBtHk0lTH72S/uZffOGeL/3gxR3u/vV/+cmpflwV3SQKffbO12fqlvTYlG3/ds3RlYqmomkWoqggS8Ik60ZJSOSOHvzA+6v/urFR8Kdf/P5PTvX/U2NPH9f8nOEvrnS/qanKE6rEIkhBKYw/CaVb5tnlOImedn3/N84t7nxveT3nn7/1+tuwbtjJ2w+Og2XM4YNVmnUTXdLJ8hzP9Tn7ZmcCVN7g2f8OAAD//zp0+rQUu87sAAAAAElFTkSuQmCC"
      },
      {
         "code": "NaM",
         "source": "BetterTTV Global Emotes",
         "image_type": "png",
         "image_data": "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"
      },
      {
         "code": "monkaS",
         "source": "BetterTTV Global Emotes",
         "image_type": "png",
         "image_data": "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"
      },
      {
         "code": ":)",
         "source": "Twitch Global",
         "image_type": "png",
         "image_data": "iVBORw0KGgoAAAANSUhEUgAAABgAAAASCAMAAAB2Mu6sAAABMlBMVEX///8AAAC6o+Pu7u6GhoaGhoaGhoaGhoaGhoaGhoa6o+O6o+O6o+O6o+O6o+O6o+O6o+O6o+OGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhoaGhobAwMDo6Oju7u7BwcHp6el9cpDt7e2pqakWFhZ9c5BkQaVxWpodHR2np6cAAAAcHBxwWprl5eVGRkYHBwcICAhFRUXi4uLr6+umpqZXV1dZWVne3t7n5+eHh4fT09PHx8ff39/j4+Pa2trV1dXd3d3Y2NjCwsK/v79mZmaBgYG4uLjLy8u2trZ9fX1eXl7k5OS0tLTq6urb29t0dHQiIiICAgIBAQEgICBxcXHU1NSzs7Pc3NzR0dHKysrPz89rDqPzAAAAJnRSTlMAAAAAJa3xJ+/zAwcKExwmFizt64B+B7m7BgmfnFDbTgRXs7JVAybtrqMAAAEqSURBVHicVJB3V/IwFIdDXgq8tGW5UXFfGkqKRZQ6AG0pUIgLleGW8f2/gicFKjz/PCf5nXvuQC5+IQBTAoIfeQRDaYVMUdKhoBcIGTJHRvCCgDIfqP99/DMsSiKQBUCUxDCS5EgWqJbLaZSQieE4IktIjup5UE8Kp2cqIRNDXo/KKBbXi2CcX1xelcrlErcBRT0e4xVZMCrXNxXTskxuA7K8QpIjVVDtWs2uNxp1bhWqvIc7FXWaTadlmi1uOpkKY4yBMHbLXO4YI+DDGLsLJgzHvn94bLefnl9sp7Pkm22+3O3Z/de394/PL+e797PinWR1rUv7fYUQZTCg3fWNv/MmNwvDDhuNWGdY2UqiebZT2tiyxlpq5x9aZHdvP5E4ODyavX8DAAD//5JsPIlnHXeLAAAAAElFTkSuQmCC"
      },
      {
         "code": "richkidHype",
         "source": "Channel: Richwcampbell",
         "image_type": "png",
         "image_data": "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"
      }
   ],
   "messages": [
      {
         "content_offset": 66.314,
         "display_name": "StreamElements",
         "display_color": "#5B99FF",
         "content": "richwcampbell is now live! Streaming Just Chatting: IM BACK--WE GOT A LOT TO TALK ABOUT",
         "badges": [
            {
               "code": "moderator",
               "version": "1"
            },
            {
               "code": "partner",
               "version": "1"
            }
         ]
      },
      {
         "content_offset": 81.333,
         "display_name": "burdefuge",
         "content": "asmon is still streaming sorry dude"
      },
      {
         "content_offset": 84.395,
         "display_name": "elilego",
         "content": "Yessssssssirrrrrrrrr",
         "badges": [
            {
               "code": "premium",
               "version": "1"
            }
         ]
      },
      {
         "content_offset": 84.907,
         "display_name": "timber_0304",
         "display_color": "#BBF87B",
         "content": "yo",
         "badges": [
            {
               "code": "glhf-pledge",
               "version": "1"
            }
         ]
      },
      {
         "content_offset": 86.344,
         "display_name": "AmbTTV",
         "display_color": "#FF0000",
         "content": "pogU",
         "badges": [
            {
               "code": "subscriber",
               "version": "6"
            }
         ]
      },
      {
         "content_offset": 89.049,
         "display_name": "Flak_",
         "display_color": "#008B50",
         "content": "hey what up, Rich",
         "badges": [
            {
               "code": "premium",
               "version": "1"
            }
         ]
      },
      {
         "content_offset": 92.126,
         "display_name": "AmbTTV",
         "display_color": "#FF0000",
         "content": "OUR STREAMER IS BACK",
         "badges": [
            {
               "code": "subscriber",
               "version": "6"
            }
         ]
      }
   ]
}
```

</p>
</details>

### CLI Arguments
```
  -client_id string
        Twitch Developer Application Client ID
  -exclude value
        What to exclude from the archival. Possible values "messages", "badges" or "emotes"
  -output_dir string
        Where to output archive files (default ".")
  -remove_unused
        Remove dangling emotes / badges which are not used throughout the entire VOD (default true)
  -vod_id string
        Target VOD ID to archive
```

### Unsupported
- Twitch Turbo variations of the Robot emotes are not archived.
- Channel Subscription Emotes from channels outside of the VOD are not archived.

### TODO:
- Investigate support for archival of subscription emotes from other channels