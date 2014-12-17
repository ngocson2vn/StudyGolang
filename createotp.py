import pyotp
import time
import base64
import hashlib

from hotpie import HOTP, TOTP

key = 'ngocson2vn@gmail.comHDECHALLENGE003'
t1 = TOTP(key, digits=10)      # <time-based-value>
print t1

secret = base64.b32encode('ngocson2vn@gmail.comHDECHALLENGE003')
totp = pyotp.TOTP(secret, 10, hashlib.sha512)
t2 = totp.now()

print t2
text = 'ngocson2vn@gmail.com:%s' % t1
print text
print base64.b64encode(text)
