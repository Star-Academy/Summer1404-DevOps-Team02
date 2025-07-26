<div dir="rtl">

Feel free to answer in Persian
- Name some of Linux popular distributions:
    - `Debian`
    - `Arch`
    - `Redhat`
    - ...
- What is the /proc directory is used for?  
  `اطلاعات مربوط به پراسس ها از قبیل namespace ها و stdin, stdout`
- Where are system-wide configuration files are stored in linux directory hierarchy?  
  `/etc`
- How linux shell commands are executed?  
  `با اجرا کردن کد باینری مربوط به هر کامند`
- Name some popular linux commands and their respective usage:  
    - `cd: برای عوض کردن دیرکتوری کنونی`
    - `ls: لیست فایل ها و دیرکتوری ها`
    - `pwd: print working directory`
    - ...
- What does this command do?  
  `:(){:|:&};:`  
  `تابعی که بصورت بازگشتی خودش را صدا میزند. ورودی مرحله n ام تابع برابر با خروجی مرحله n-1 ام تابع است. این کد باعث پر شدن memory میشود`
- How new packages are installed in linux?  
  `با سرچ کردن در ریپازیتوری های تعیین شده پکیج را پیدا کرده و سپس دانلود میکند`
- Given the following `ls` command output, if danny is only inside the group danny, what files can he read?  
    ```
    drwxr-xr-x  2 danny danny 4.0K May 25 23:06 .
    drwxr-x--- 25 danny danny 4.0K Jul  8 22:19 ..
    -rwxrwxrwx  1 danny admin    0 May 25 23:02 a
    ----r--r--  1 danny games    0 May 25 23:02 b
    -r--rw----  1 root  root     0 May 25 23:02 c
    -r--r-----  1 root  danny    0 May 25 23:02 d
    ```
    `1, 2, 3, 6`
- What is `.` file and what happens if its permissions are set to `000`?  
  `همان current directory است. پس از تغییر دادن پرمیشن ها دیگر نمیتوانیم داخل آن  ls بزنیم. همچنین در صورتی که از این دیرکتوری بیرون بیاییم دیگر نمیتوانیم وارد آن شویم و فقط با دسترسی یوزر root میتوانیم این کار ها را بکنیم `
- How can you exit Vim?  
  `اول بایستی وارد command mode شویم و سپس با دستور :q میتوانیم از آن خارج شویم. البته اگر تغییری در فایل داده باشیم بایستی از :wq برای سیو کردن فایل یا :!q برای خارج شدن بصورت force استفاده کنیم`
- Who is the murderer running around in terminal city? How did you find them?  
  `Jeremy Bowers`
## Review
Link to your PR:  
`[LINK TO YOUR PR]`  
 - [ ] Your PR is reviewed and approved by both mentors
 - [ ] Your PR is merged

</div>
