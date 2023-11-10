'use client'

import React ,{useState,FormEvent}from 'react';
import Auth from '../components/Auth';
import { authUrl } from '@/utils/network';
import axios, { AxiosError, AxiosResponse } from 'axios';
import { errorHandler } from '@/utils/errorHandler';
import { useRouter } from 'next/navigation';
import { userTokenKey } from '@/utils/constants';

import useAxiosHandler from '@/utils/axiosHandler';
import withoutAuth from '../hocs/withoutAuth';

interface LoginType {
  token: string
}

const Login:React.FC = () => {
    const router = useRouter()
    const [loading, SetLoading] = useState(false);
    const {axiosHandler} = useAxiosHandler()
    const onSubmit =async(
        e: FormEvent<HTMLFormElement>,
        formRef: React.RefObject<HTMLFormElement>
    )=>{
    e.preventDefault();
    SetLoading(true);
    let arg = {
      email: formRef.current?.email.value,
      password: formRef.current?.password.value,
    };

    const response = await axiosHandler<LoginType>({
      method: "POST",
      url: authUrl.login,
      data: arg
    })
    
     SetLoading(false);
  
  

    if (response.data) {
      localStorage.setItem(userTokenKey, response.data.token)
        router.push("/accounting");
      }


    }
    
    return  <Auth
     loading={loading} 
     showRemembered
     title="Log In"
     buttonTitle = "Login"
     onSubmit={onSubmit} />;
}
export default withoutAuth(Login);


