import { Modal, Typography } from 'antd';
import { RcFile } from 'antd/es/upload';
import React, { useEffect, useState } from 'react';

import { AvatarUploadField } from '@/components/Avatar';
import { UserStatus } from '@/services/fuck-web/enums';
import { uploadFile as postFile } from '@/services/fuck-web/files';
import { getRoles } from '@/services/fuck-web/roles';
import { IntlContext } from '@/utils/intl';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import ProForm, { ModalForm, ProFormSelect, ProFormText, StepsForm } from '@ant-design/pro-form';

export type FormValueType = Omit<API.UpdateUserRequest, 'id'> & {
  id?: string;
};
export type UpdateFormProps = {
  onCancel: (flag?: boolean, formVals?: FormValueType) => void;
  onSubmit: (values: FormValueType) => Promise<boolean>;
  modalVisible: boolean;
  values?: API.UserInfo;
  title?: React.ReactNode;
  parentIntl: IntlContext;
};

const CreateOrUpdateForm: React.FC<UpdateFormProps> = ({ parentIntl, ...props }) => {
  const intl = new IntlContext('form', parentIntl);
  const { values: dfValues, title, modalVisible, onSubmit, onCancel } = props;

  return (
    <ModalForm<FormValueType>
      // formProps={{
      //   preserve: false,
      // }}
      labelCol={{ span: 8 }}
      wrapperCol={{ span: 14 }}
      layout={'horizontal'}
      title={intl.t('title.basicConfig', 'Basic')}
      modalProps={{
        width: 640,
        bodyStyle: { padding: '32px 40px 48px' },
        destroyOnClose: true,
        title: title,
        onCancel: () => {
          Modal.confirm({
            title: intl.t('cancel?', 'Cancel editing?'),
            icon: <ExclamationCircleOutlined />,
            onOk() {
              onCancel();
            },
            maskClosable: true,
          });
        },
      }}
      open={modalVisible}
      initialValues={{ ...dfValues }}
      onFinish={async (values) => {
        return onSubmit({
          ...values,
          status: dfValues?.status ? dfValues.status : UserStatus.normal,
          isDelete: dfValues?.isDelete ? dfValues?.isDelete : false,
        });
      }}
    >
      <AvatarUploadField
        label={intl.t('avatar.label', 'Avatar')}
        name={'avatar'}
        request={async (filename: string, fileObj: RcFile | string | Blob) => {
          const formData = new FormData();
          formData.append(filename, fileObj);
          const resp = await postFile({ data: formData, requestType: 'form' });
          if (resp.data) {
            return resp.data[filename];
          }
          return '';
        }}
      />
      <ProFormText hidden={true} name="id" />
      <ProFormText
        name="username"
        label={intl.t('userName.label', 'Username')}
        width="md"
        rules={[
          {
            required: true,
            message: intl.t('userName.required', 'Please input username!'),
          },
          {
            pattern: /^[-_A-Za-z0-9]+$/,
            message: intl.t('name.invalid', 'Username format error!'),
          },
        ]}
      />

      <ProFormText name="fullName" label={intl.t('fullName.label', 'FullName')} width="md" />
      <ProFormText name="email" label={intl.t('email.label', 'Email')} width="md" />
      <ProFormText
        name="phoneNumber"
        label={intl.t('phoneNumber.label', 'Telephone number')}
        width="md"
      />
      <ProFormSelect
        name="roleId"
        label={intl.t('role.label', 'Role')}
        width="md"
        fieldProps={{ optionLabelProp: 'name' }}
        request={async () => {
          const roles = (await getRoles({})).data ?? [];
          return roles.map(({ id, name, describe }) => {
            return {
              label: (
                <div>
                  {name}
                  <Typography.Text
                    style={{ width: 200, color: '#827d7d', marginLeft: 10 }}
                    ellipsis
                  >
                    {describe}
                  </Typography.Text>
                </div>
              ),
              value: id,
              name,
            };
          });
        }}
      />
    </ModalForm>
  );
};

export default CreateOrUpdateForm;
