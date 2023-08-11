import { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';
import { CreateMsgSubscriberInput, MsgType } from '@/__generated__/msgsrv/graphql';
import { DrawerForm } from '@ant-design/pro-components';
import { createSub, delSub, getMsgTypeAndSubInfo } from '@/services/msgsrv/type';
import { Radio, Skeleton, Space, Transfer, Typography, message } from 'antd';
import { getOrgUserList } from '@/services/adminx/org/user';
import store from '@/store';
import { getOrgGroupList } from '@/services/adminx/org/role';
import { TransferItem } from 'antd/es/transfer';
import { useLeavePrompt } from '@knockout-js/layout';

type SubjectType = 'user' | 'exUser' | 'orgRole'

export default (props: {
  open: boolean;
  title?: string;
  id: string;
  onClose: (isSuccess?: boolean) => void;
}) => {
  const { t } = useTranslation(),
    [userState] = store.useModel('user'),
    [subject, setSubject] = useState<SubjectType>('user'),
    [, setLeavePromptWhen] = useLeavePrompt(),
    [info, setInfo] = useState<MsgType>(),
    [loading, setLoading] = useState(false),
    [dataSource, setDataSource] = useState<TransferItem[]>([]),
    [targetKeys, setTargetKeys] = useState<string[]>([]),
    [saveLoading, setSaveLoading] = useState(false),
    [saveDisabled, setSaveDisabled] = useState(true);

  useEffect(() => {
    setLeavePromptWhen(saveDisabled);
  }, [saveDisabled]);

  const
    onOpenChange = (open: boolean) => {
      if (!open) {
        props.onClose?.();
      }
      setSaveDisabled(true);
    },
    getRequest = async () => {
      setLoading(true);
      const result = await getMsgTypeAndSubInfo(props.id);
      if (result?.id) {
        setInfo(result as MsgType);
      }
      setLoading(false);
      return {}
    },
    getTransferData = async () => {
      const data: TransferItem[] = [], tkeys: string[] = [];
      if ('user' === subject) {
        const result = await getOrgUserList(userState.tenantId, {
          current: 1,
          pageSize: 999,
        })
        if (result?.totalCount) {
          result.edges?.forEach(item => {
            if (item?.node) {
              data.push({
                key: item.node.id,
                name: item.node.displayName,
                description: item.node.email || "",
              })
            }
          })
        }
        info?.subscriberUsers?.forEach(item => {
          if (item.userID) {
            tkeys.push(item.userID)
          }
        })
      } else if ('exUser' === subject) {
        const result = await getOrgUserList(userState.tenantId, {
          current: 1,
          pageSize: 999,
        })
        if (result?.totalCount) {
          result.edges?.forEach(item => {
            if (item?.node) {
              data.push({
                key: item.node.id,
                name: item.node.displayName,
                description: item.node.email || "",
              })
            }
          })
        }
        info?.excludeSubscriberUsers?.forEach(item => {
          if (item.userID) {
            tkeys.push(item.userID)
          }
        })
      } else if ('orgRole' === subject) {
        const result = await getOrgGroupList({
          current: 1,
          pageSize: 999,
        })
        if (result?.totalCount) {
          result.edges?.forEach(item => {
            if (item?.node) {
              data.push({
                key: item.node.id,
                name: item.node.name,
                description: item.node.comments || '',
              })
            }
          })
        }
        info?.subscriberRoles?.forEach(item => {
          if (item.orgRoleID) {
            tkeys.push(`${item.orgRoleID}`);
          }
        })
      }
      setTargetKeys(tkeys);
      setDataSource(data);
      setSaveDisabled(true);
    },
    onFinish = async () => {
      let isTrue = false;
      if (info) {
        const oldKyes = info?.subscriberUsers?.map(item => item.userID as string) || [],
          addKeys = targetKeys.filter(key => !oldKyes.includes(key)),
          delKeys = oldKyes.filter(key => !targetKeys.includes(key));
        setSaveLoading(true);
        if (addKeys.length) {
          const inputs = addKeys.map(key => {
            const data: CreateMsgSubscriberInput = {
              msgTypeID: info.id,
              tenantID: userState.tenantId,
            }
            if (subject === 'user') {
              data.userID = key
              data.exclude = false
            } else if (subject === 'exUser') {
              data.userID = key
              data.exclude = true
            } else if (subject === 'orgRole') {
              data.orgRoleID = key
              data.exclude = false
            }
            return data;
          });

          const result = await createSub(inputs);
          if (result?.length) {
            isTrue = true;
          }
        }
        if (delKeys.length) {
          const ids = delKeys.map(key => {
            let id: string = '';
            if (subject === 'user') {
              id = info.subscriberUsers.find(item => item.userID == key)?.id as string
            } else if (subject === 'exUser') {
              id = info.excludeSubscriberUsers.find(item => item.userID == key)?.id as string
            } else if (subject === 'orgRole') {
              id = info.subscriberRoles.find(item => `${item.orgRoleID}` == key)?.id as string
            }
            return id
          })
          const result = await delSub(ids);
          if (result) {
            isTrue = true;
          }
        }

        if (isTrue) {
          message.success(t('submit_success'))
        }
        props.onClose?.(isTrue);
        setSaveLoading(false);
      }
      return false;
    }

  useEffect(() => {
    if (info) {
      getTransferData();
    }
  }, [subject, info])

  return (
    <DrawerForm
      title={props.title}
      open={props.open}
      submitter={{
        searchConfig: {
          submitText: t('submit'),
          resetText: t('cancel'),
        },
        submitButtonProps: {
          loading: saveLoading,
          disabled: saveDisabled,
        },
      }}
      drawerProps={{
        width: 600,
        destroyOnClose: true,
      }}
      request={getRequest}
      onFinish={onFinish}
      onOpenChange={onOpenChange}
    >
      {
        loading ? <Skeleton /> : info ?
          <Space direction="vertical">
            <div>
              {t('msg_type')}：{info.category}-{info.name}
            </div>
            <div>
              {t('receiving_subject')}：<Radio.Group
                value={subject}
                onChange={(e) => {
                  setSubject(e.target.value)
                }} >
                <Radio.Button value="user">{t('receiving_user')}</Radio.Button>
                <Radio.Button value="orgRole">{t('receiving_user_group')}</Radio.Button>
                <Radio.Button value="exUser">{t('exclude_user')}</Radio.Button>
              </Radio.Group>
            </div>
            <div>
              <Transfer
                oneWay
                dataSource={dataSource}
                targetKeys={targetKeys}
                onChange={(newKyes) => {
                  setSaveDisabled(false);
                  setTargetKeys(newKyes);
                }}
                render={(item) => {
                  return <div key={item.id}>
                    <div>{item.name}</div>
                    <Typography.Text type="secondary">
                      {item.description}
                    </Typography.Text>
                  </div>
                }}
                listStyle={{ height: '70vh' }}
                pagination={{
                  pageSize: 15,
                }}
              />
            </div>
          </Space>
          : <></>
      }
    </DrawerForm>
  );
};
