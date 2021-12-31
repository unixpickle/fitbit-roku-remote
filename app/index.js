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
});

function setEnabled(flag) {
  const opacity = flag ? 1 : 0.5;
  Object.keys(buttonIdToKey).forEach((buttonId) => {
    const button = document.getElementById(buttonId);
    button.style.opacity = opacity;
  });
}

// Show enabled state based on connection to companion app.
setEnabled(false);
messaging.peerSocket.addEventListener("open", () => setEnabled(true));
messaging.peerSocket.addEventListener("close", () => setEnabled(false));
