import Editor from '@monaco-editor/react';
import { Drawer } from 'antd';


export default (props: {
  open: boolean;
  title?: string;
  onClose: () => void;
}) => {
  const text =
    `# the receiver name
name: 'team-X-mails'
# the email config
emailConfigs:
    # The email address to send notifications to
  - to: 'team-X+alerts@example.org'
    # The sender address
    from: "xxx@qq.com"
    # The SMTP host and port through which emails are sent. E.g. example.com:25
    smartHost: "smtp.qq.com:587"
    # Further headers email header key/value pairs.
    # Overrides any headers previously set by the notification implementation
    headers:
    # sets the auth type of the SMTP client
    # the vaule is "" or "PLAIN" or "LOGIN" or "CRAM-MD5"
    authType: "PLAIN"
    # The username to use for authentication
    authUsername: "xxx@qq.com"
    # The password to use for authentication
    authPassword: "xxx"
    # The identity to use for authentication
    authIdentity: ""
    # the CRAM-MD5 secret
    authSecret: ""
messageConfig:
  to: ""
  subject: ""
  redirect: ""
    `;


  return (
    <Drawer
      width={800}
      title={props.title}
      open={props?.open}
      maskClosable={false}
      onClose={() => {
        props.onClose?.();
      }}
    >
      <Editor
        className="adminx-editor"
        height="80vh"
        value={text}
        options={{
          readOnly: true,
        }}
        defaultLanguage="yaml"
      />
    </Drawer>
  );
};
