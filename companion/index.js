import * as messaging from "messaging";
import { settingsStorage } from "settings";

messaging.peerSocket.addEventListener("message", (evt) => {
  const key = evt.data['key'];
  const url = getECPURL() + '/keypress/' + key;
  fetch(url, { method: 'POST' });
});

function getECPURL() {
  return 'http://' + getECPHost() + ":" + getECPPort();
}

function getECPHost() {
  return JSON.parse(settingsStorage.getItem('ecphost') || '{"name":"192.168.0.211"}')["name"];
}

function getECPPort() {
  return parseInt(JSON.parse(settingsStorage.getItem('ecphost') || '{"name":"8060"}')["name"]);
}
