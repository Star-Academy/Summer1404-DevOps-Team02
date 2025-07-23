Feel free to answer in Persian
- What is a protocol?
  `قرار دادی است که با استفاده از آن دیتایی را به مکانی دیگر انتقال میدهیم`
- Why computer networks architecture is layered?  
  `برای توسعه پذیر شدن هر لایه. در غیر اینصورت پیچیدگی بالا میرفت و با لایه ای بودن، هر قسمت وظیفه مشخص خود را دارد`
- How your national code and a computer's MAC address are simillar?
  `هر دو یونیک و فیزیکی هستند و روشی برای شناسایی`
- How ARP works? Why do we need it?
  `ARP is Access resolution protocol. 
وقتی سیستمی وارد شبکه میشود یک جدول ARP ساخته میشود که هر IP را به مک آدرس مورد نظر نظیر میکند
چون ethernet frame از مک آدرس استفاده میکند
`
- How computers that are not in a subnet comminucate with each other? explain.
  `پس از آنکه پکت به سمت روتر همان subnet رفت، وظیفه روتر این است که با ارتباط گرفتن با بقیه روتر ها با استفاده از ip تصمیم بگیرد که پکت را به کجا بفرستد.`
- What is the importance of TCP and why a two-way handshake won't work for it?
  `چون tcp از ارسال پکت مطمعن هست به صورت سالم 
اگر two way handshake استفاد شود. نمیتوان sequence number ها را سینک کنیم و reliable نیست
`
- What is the TTL field in a DNS record?
  `مدت زمانی که یک فیلد در کش dns record میماند`
- How to open a TCP port and connect to it using Netcat?
  `nc -l host port
   nc host port 
`
