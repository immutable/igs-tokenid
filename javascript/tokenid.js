const { uuidv7 } = require("uuidv7");

const uuid = uuidv7().replace(/-/g, '');
console.log('uuid', uuid);

const tokenid = BigInt(`0x${uuid}`, 16);
console.log('tokenid', tokenid.toString());