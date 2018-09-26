# _*_ -coding:utf8-


# pow(x,y,z) and the z "(x**y)%z"
# max :: 1.single args 2. iter 
# round(x,y) :: 
# use tsinghua yuan install pip
# pip install xxx -i https://pypi.tuna.tsinghua.xxxxm, 
# math.log :: e  
# math.log(10,10)
# math.sin()  input huduzhi
# EP1:
'''
import math
:
x1, y1 = eval(input('输入A点坐标：'))
x2, y2 = eval(input('输入B点坐标：'))
x3, y3 = eval(input('输入C点坐标：'))

a = math.sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2)
b = math.sqrt((x1 - x3) ** 2 + (y1 - y3) ** 2)
c = math.sqrt((x2 - x3) ** 2 + (y2 - y3) ** 2)

A = math.degrees(math.acos((a * a - b * b - c * c) / (-2 * b * c)))
B = math.degrees(math.acos((b * b - a * a - c * c) / (-2 * a * c)))
C = math.degrees(math.acos((c * c - b * b - a * a) / (-2 * a * b)))

print('三角形的三个角分别为', A, B, C)

     
'''
a = '''
ajshbdfkajs
askhfbajsf
ashbfkajsb
aksfjbkj
'''
# 
# print(type(a))
# ord() input string, return ascii value
# chr()
# from __future__ import print_function
# loading  future package in print function
res = ''
for i in 'maomaochong@163.com':
    res = res +  chr(ord(i) + 1)
    # print(res,end='')
print(res)
