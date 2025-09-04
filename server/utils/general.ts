// 16位用户注册唯一ID
function generateUserId() {
  const prefix = "U";
  const timestamp = Date.now().toString();
  const rand = Math.floor(Math.random() * 90 + 10).toString(); 
  return prefix + timestamp + rand;
}


export {generateUserId}