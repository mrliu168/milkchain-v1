Step 2:
  ```js
  npm install
  //安装时速度可能会很慢，静等即可
  ```
  
Step 3:
  ```
  ./startFabric.sh
  若遇到权限问题执行chmod a+x startFabric.sh
  若仍有问题进入basic-network文件夹下执行 chmod a+x start.sh
  ```

 Step 4:
   ```
   node registerAdmin.js
   node registerUser.js
   node server.js
   ```

访问`http://localhost:7000`
