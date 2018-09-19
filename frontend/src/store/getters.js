console.log('*** inside the store/getters file ***');

export const testString = state => state.testString;
export const localJWT = state => state.localJWT;
export const siteName = state => state.siteName;
export const loggedIn = state => state.loggedIn;
export const email = state => state.email;
export const channelData = state => state.channelData;
export const subChannelData = state => state.subChannelData;
export const profileMsgArray = state => state.profileMsgArray;
export const profileCanvas = state => state.profileCanvas;
export const usersList = state => state.usersList;
export const receivedMail = state => state.receivedMail;
export const sentMail = state => state.sentMail;
export const msgSentTime = state => state.msgSentTime;