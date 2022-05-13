# Go-RLT
A rate limit testing tool written in GO

Recently I found the need to test if a server had rate limiting enabled or not? I decided that this was a pretty good opportunity to learn GO! 

So I wrote a really basic program in GO to do just that. It was able to make over 1000 requests in around 1 second! and that was without any threading! Just a simple for loop. Meanwhile, my good ol' Python program is still running as we speak. 

Update:

I created a web server in golang to match the times required by Python and Go to make 10,000 requests. I used threading in python and Go routines in go. Here are the results.

![](https://i.imgur.com/H6PQB3H.png)

PS: Those who may say that this is an unfair comparison, or both the languages have their own benefits. I agree. I just did this because it was fun.