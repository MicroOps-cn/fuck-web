import { Space } from 'antd';
import React, { useEffect, useState } from 'react';

import { QuestionCircleOutlined, HomeOutlined } from '@ant-design/icons';
import { useModel } from '@umijs/max';

import SelectLang from '../SelectLang';
import Avatar from './AvatarDropdown';
import styles from './index.less';

export const getActions = (role?: string) => {
  // const [externalUrl, setExternalUrl] = useState<string>('');
  // useEffect(() => {
  //   const { initialState } = useModel('@@initialState');
  //   if (initialState?.globalConfig?.external_url) {
  //     setExternalUrl(initialState.globalConfig.external_url);
  //   }
  // });
  return [
    <div className={styles.action} title={'Help'} key={'help'}>
      <QuestionCircleOutlined />
    </div>,
    <Avatar menu={true} key={'avatar'} />,
    //{/* <NoticeIconView /> */}
    <SelectLang className={styles.action} key={'lang'} />,
  ];
};

const GlobalHeaderRight: React.FC = () => {
  const { initialState } = useModel('@@initialState');
  if (!initialState || !initialState.settings) {
    return null;
  }
  const { role } = initialState.currentUser ?? {};
  return (
    <Space className={styles.right}>
      <a
        className={styles.action}
        style={{ color: 'unset', padding: '0 12px' }}
        hidden={!initialState?.currentUser?.role}
        href={initialState?.globalConfig?.external_url ?? '/'}
        title={'Dashboard'}
        key={'admin'}
      >
        <HomeOutlined />
      </a>
      {getActions(role)}
    </Space>
  );
};
export default GlobalHeaderRight;
