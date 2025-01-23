'use client';

import { useState } from 'react';
import { useForm } from 'react-hook-form';
import axios from 'axios';

const LoginPage = () => {
  const { register, handleSubmit, formState: { errors } } = useForm();
  const [errorMessage, setErrorMessage] = useState('');

  const onSubmit = async (data: any) => {
    try {
      const response = await axios.post('APIエンドポイント', data);
      // 成功時の処理（例: トークンを保存）
      console.log(response.data);
    } catch (error) {
      setErrorMessage('ログインに失敗しました。再度お試しください。');
    }
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-100">
      <form onSubmit={handleSubmit(onSubmit)} className="bg-white p-8 rounded-lg shadow-lg w-96">
        <h2 className="text-2xl font-bold mb-6 text-center">ログイン</h2>

        {errorMessage && <div className="text-red-500 mb-4">{errorMessage}</div>}

        <div className="mb-4">
          <label htmlFor="username" className="block text-sm font-semibold">ユーザー名</label>
          <input
            id="username"
            type="text"
            {...register('username', { required: 'ユーザー名は必須です' })}
            className="w-full p-2 border border-gray-300 rounded mt-2"
          />
          {errors.username && typeof errors.username.message === 'string' && (
  <p className="text-red-500 text-xs">{errors.username.message}</p>
)}
        </div>

        <div className="mb-6">
          <label htmlFor="password" className="block text-sm font-semibold">パスワード</label>｀｀｀｀
          <input
            id="password"
            type="password"
            {...register('password', { required: 'パスワードは必須です' })}
            className="w-full p-2 border border-gray-300 rounded mt-2"
          />
          {errors.password && typeof errors.password.message === 'string' && (
  <p className="text-red-500 text-xs">{errors.password.message}</p>
)}
        </div>

        <button type="submit" className="w-full py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
          ログイン
        </button>
      </form>
    </div>
  );
};

export default LoginPage;
