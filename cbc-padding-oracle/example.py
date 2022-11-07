#!/usr/bin/env python3

# CBC padding oracle attack
# - lenerd

import requests
import sys

def test_systems_security(base_url):
    res = requests.get(f'{base_url}/', verify="ca.pem")
    ciphertext = bytes.fromhex(res.cookies['authtoken'])
    print(f'[+] received ciphertext: {ciphertext.hex()}')
    res = requests.get(f'{base_url}/check/', cookies={'authtoken': ciphertext.hex()}, verify="ca.pem")
    print(f'[+] done:\n{res.text}')

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print(f'usage: {sys.argv[0]} <base url>', file=sys.stderr)
        exit(1)
    test_systems_security(sys.argv[1])
