import { Button, message } from 'antd';
import type { StoreValue } from 'antd/lib/form/interface';
import React, { useState } from 'react';
import { history, useIntl, SelectLang, useModel } from 'umi';

import defaultSettings from '@/../config/defaultSettings';
import { loginPath } from '@/../config/env';
import Footer from '@/components/Footer';
import { resetPassword } from '@/services/fuck-web/user';
import { IntlContext } from '@/utils/intl';
import { getPublicPath } from '@/utils/request';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { ProFormText, LoginForm } from '@ant-design/pro-form';

import styles from './index.less';

const ResetPassword: React.FC<Record<string, any>> = (props) => {
  const rootIntl = useIntl();
  const intl = new IntlContext('pages.resetPassword', rootIntl);
  const handleResetPassword = async (values: API.ResetUserPasswordRequest): Promise<boolean> => {
    try {
      // 登录
      const msg = await resetPassword({ ...values });
      if (msg.success) {
        const defaultSuccessMessage = intl.t(
          'success',
          'The password is reset successfully. Please login again with the new password.',
        );
        message.success(defaultSuccessMessage);
        return true;
      }
    } catch (error) {
      console.error(error);
    }
    return false;
  };
  const query = props.location.query;
  const [loading, setLoading] = useState<boolean>(false);
  const { initialState } = useModel('@@initialState');
  const globalConfig = initialState?.globalConfig ?? null;
  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <LoginForm<API.ResetUserPasswordRequest>
          logo={globalConfig?.logo ?? getPublicPath('logo.svg')}
          title={globalConfig?.title ?? defaultSettings.title}
          subTitle={<> </>}
          initialValues={{
            username: query.username,
            token: query.token,
          }}
          submitter={{
            render: (submitProps) => {
              return (
                <Button loading={loading} onClick={submitProps.submit} block type="primary">
                  重置密码
                </Button>
              );
            },
          }}
          onFinish={async (values) => {
            setLoading(true);
            if (
              await handleResetPassword({
                newPassword: values.newPassword,
                oldPassword: values.oldPassword,
                userId: query.userId,
                token: query.token,
                username: query.username,
              })
            ) {
              history.push(loginPath);
            }
            setLoading(false);
          }}
        >
          <ProFormText
            fieldProps={{
              value: query.username,
              size: 'large',
              disabled: true,
              prefix: <UserOutlined className={styles.prefixIcon} />,
            }}
          />
          <ProFormText.Password
            name="oldPassword"
            fieldProps={{
              size: 'large',
              prefix: <LockOutlined className={styles.prefixIcon} />,
            }}
            hidden={query.token}
            placeholder={intl.t('oldPassword.placeholder', 'Please enter current password')}
            rules={[
              {
                required: !query.token,
                message: intl.t('oldPassword.required', 'Please enter current password!'),
              },
            ]}
          />
          <ProFormText.Password
            name="newPassword"
            fieldProps={{
              size: 'large',
              prefix: <LockOutlined className={styles.prefixIcon} />,
            }}
            placeholder={intl.t('password.placeholder', 'Please enter a new password')}
            rules={[
              {
                required: true,
                message: intl.t('password.required', 'Please enter a new password!'),
              },
            ]}
          />
          <ProFormText.Password
            name="newPasswordConfirm"
            fieldProps={{
              size: 'large',
              prefix: <LockOutlined className={styles.prefixIcon} />,
            }}
            placeholder={intl.t('confirmPassword.placeholder', 'Confirm new password.')}
            rules={[
              {
                required: true,
                message: intl.t('confirmPassword.required', 'Confirm new password!'),
              },
              ({ getFieldValue }) => ({
                validator: (_, value: StoreValue) => {
                  if (!value || getFieldValue('newPassword') === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(
                    new Error('The two passwords that you entered do not match!'),
                  );
                },
              }),
            ]}
          />
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default ResetPassword;
