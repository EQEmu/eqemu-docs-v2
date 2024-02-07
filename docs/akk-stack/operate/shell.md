!!! info
    Now that you're installed we need to look at how we interact with the environment.

    There are many ways to shell into services in the stack, this is a quick overview of the most common methods


### Direct Bash

From the host level, you can use **make bash** to hop into the **eqemu-server** container

<img src="https://user-images.githubusercontent.com/3319450/87241544-e8473b00-c3e9-11ea-8232-33fa3da9d40b.gif" style="border-radius: 15px">

### SSH

You can use the default SSH port **2222** to shell into the **eqemu-server** container remotely. This password is generated at install time and is persistent through reboots and can be found in the **.env** file

<img src="https://user-images.githubusercontent.com/3319450/87241545-ea10fe80-c3e9-11ea-9a7f-c97ba54e93fa.gif" style="border-radius: 15px">

### MySQL Console

You can obtain a MySQL shell from either **eqemu-server** **make mc** or from the **eqemu-server** embedded shell alias **mc**

<img src="https://user-images.githubusercontent.com/3319450/87241546-ec735880-c3e9-11ea-9a8e-412ca4d99736.gif" style="border-radius: 15px">