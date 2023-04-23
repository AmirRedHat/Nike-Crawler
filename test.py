import json 

f = "./test.json"

_file = open(f, "r")

x = json.loads(_file.read())
x = json.loads(x)

# ["products"]["DM0032-800"]["skus"]
print(x["props"]["pageProps"]["initialState"]["Threads"])

# def formatData(t,s):
#     if not isinstance(t,dict) and not isinstance(t,list):
#         print("\t"*s+str(t))
#     else:
#         for key in t:
#             print("\t"*s+str(key))
#             if not isinstance(t,list):
#                 formatData(t[key],s+1)

# print("working")
# def findPath(x):
#    for k, v in x.items() :
#        print(k)
#        print(type(v))
#        if "vas" == k:
#            print("found")
#            return v
#        else:
#            if type(v) == dict:
#             return findPath(v)
           
# y = findPath(x)