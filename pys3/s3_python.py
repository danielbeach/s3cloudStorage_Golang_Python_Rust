import boto3
import os
from datetime import datetime

t1 = datetime.now()
os.environ['AWS_ACCESS_KEY_ID'] = ''
os.environ['AWS_SECRET_ACCESS_KEY'] = ''

s3 = boto3.resource('s3')
bucket = s3.Bucket('confessions-of-a-data-guy')
for obj in bucket.objects.filter(Prefix='202201'):
    print(obj.key)
t2 = datetime.now()
print("I took {x}".format(x=t2-t1))
