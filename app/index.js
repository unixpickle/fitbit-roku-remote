import * as document from 'document';
import * as messaging from "messaging";

function sendKey(key) {
  if (messaging.peerSocket.readyState === messaging.peerSocket.OPEN) {
    messaging.peerSocket.send({ key: key });
  }
}

const buttonIdToKey = {
  'up-button': 'Up',
  'down-button': 'Down',
  'left-button': 'Left',
  'right-button': 'Right',
  'ok-button': 'Select',
};

Object.keys(buttonIdToKey).forEach((buttonId) => {
  const button = document.getElementById(buttonId);
  button.addEventListener('click', () => {
    sendKey(buttonIdToKey[buttonId]);
  });
  button.style.display = "none";
});

// Show buttons only once we are connected to companion app.
messaging.peerSocket.addEventListener("open", () => {
  Object.keys(buttonIdToKey).forEach((buttonId) => {
    const button = document.getElementById(buttonId);
    button.style.display = "";
  });
});
