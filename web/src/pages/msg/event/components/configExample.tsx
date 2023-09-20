import Editor from '@monaco-editor/react';
import { Drawer } from 'antd';


export default (props: {
  open: boolean;
  title?: string;
  onClose: () => void;
}) => {
  const text =
    `# the receiver name
receiver: email
# List of labels to group by. Labels must not be repeated
# (unique list). Special label "..." (aggregate by all
# possible labels), if provided, must be the only element in the list
groupBy: ['alertname', 'tenant','user']
# How long to wait before sending an updated notification. E.g. "5m"
groupInterval: 5m
# How long to wait before sending the initial notification. E.g. "30s"
groupWait: 30s
# How long to wait before repeating the last notification. E.g. "4h"
repeatInterval: 4h
# MuteTimeIntervals is a list of MuteTimeInterval names
# that will mute this route when matched
muteTimeIntervals: ['']
#
activeTimeIntervals: ['']
# List of matchers that the alert’s labels should match
matchers:
    # Type to match. the value is = or != or =~ or !~
  - type: =
    # Label to match
    name: app
    # Label value to match
    value: "1"
  - type: =
    name: alertname
    value: AlterPassword
# Boolean indicating whether an alert should continue matching subsequent sibling nodes.
# It will always be overridden to true for the first-level route
continue: false

# the config like the first-level route
routes:
  - receiver: email
    matchers:
      - type: =~
        name: receiver
        value: email
    continue: true
  - receiver: internal
    matchers:
      - type: =~
        name: receiver
        value: internal
    continue: false
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
