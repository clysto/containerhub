import { FitAddon } from '@xterm/addon-fit';
import { Terminal } from '@xterm/xterm';
import { WebglAddon } from '@xterm/addon-webgl';
import { AttachAddon } from '@xterm/addon-attach';
import style from '../styles/ssh.module.css';
import { getJWTToken } from '../api';
import classNames from 'classnames';

function SSHPage() {
  let terminal = null;
  const fitAddon = new FitAddon();
  const containerID = m.route.param('id');

  window.addEventListener(
    'resize',
    () => {
      fitAddon.fit();
      sendSize();
    },
    false
  );

  const webSocket = new WebSocket(
    'ws://' + window.location.host + '/api/v1/ssh' + '?id=' + containerID + '&token=' + getJWTToken()
  );

  function sendSize() {
    const windowSize = { high: terminal.rows, width: terminal.cols };
    const blob = new Blob([JSON.stringify(windowSize)], { type: 'application/json' });
    webSocket.send(blob);
  }

  webSocket.onopen = sendSize;

  return {
    oncreate() {
      terminal = new Terminal({
        fontFamily: 'SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace',
        rows: 40,
      });
      terminal.open(document.getElementById('terminal'));
      terminal.loadAddon(fitAddon);
      terminal.loadAddon(new WebglAddon());
      const attachAddon = new AttachAddon(webSocket);
      terminal.loadAddon(attachAddon);
      fitAddon.fit();
    },
    view() {
      return (
        <div
          class={classNames(style.terminal)}
          id="terminal"
          style="height: calc(100vh); background-color: #000;"
        ></div>
      );
    },
  };
}

export default SSHPage;
