// 本地缓存服务
const PREFIX = 'ginessential_';

// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 存储
const set = (key, value) => {
    localStorage.setItem(key, value);
};

//读取
const get = (key) => {
    localStorage.getItem(key);
};

export default {
    set, get, USER_TOKEN, USER_INFO
};
