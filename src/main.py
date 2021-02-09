import socket



UDP_IP = "10.100.23.147:51832"
UDP_PORT = 20023
MESSAGE = b"yoyoyo"

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.sendto(MESSAGE,(UDP_IP, UDP_PORT))