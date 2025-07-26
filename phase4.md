Feel free to answer in Persian

- What is the importance of having a HA infrastructure?
  
  `دسترسی بالا باعث می شود که سرویس ها همیشه در دسترس باشند و در صورت وقوع حادثه یا خرابی سرویس ها و اپلیکیشن ها به ارایه سرویس به کاربران ادامه دهند که بسیار مهم است زیرا حتی چند دقیقه قطعی و عدم دسترسی به یک سرویس ممکن است شرکت را دچار ضرر مالی بسیار کند یا اعتبار آن را خدشه دار کند.`
  
- If we can afford 1s of downtime per day, what's our availabilty percantage for a year? explain.
  
  `365 ثانیه در سال می شود که باید تقسیم بر کل ثانیه های سال شود. که بالای 99.9999 است`
  
- What is failover and how can it help us achive HA?
  
  `در زمان failover منابع و درخواست ها به یک سرور جایگزین منتقل میشود. یا سرویس فیل شده اتوماتیک برمیگردد. به صورت کلی failover جلوی قطعی کامل سرویس را میگیرد.`
  
- What are some methods to ensure HA and DR in our systems?
  
  `Observing and monitoring the services`
  `Load Balancing`
  `Preventing from single point of failures`
  `Clustering and replications`
  
- Explain chaos engineering and name some tools that help us with implementing it.

  `Chaos Monkey , Gremlin are the tools
برای تست کردن system resilliency است که به صورت عمدی در سرویس ایجاد اختلال نقاط ضعف را شناسایی میکنند
`

