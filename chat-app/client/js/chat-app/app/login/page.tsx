'use client'

import { useContext, useEffect, useState } from 'react';
import {API_URL} from '../../constants'
import {useRouter } from 'next/navigation' 
import { AuthContext, UserInfo } from '@/modules/auth_provider';

const Index = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const {authenticated} = useContext(AuthContext)


  const router = useRouter()

  useEffect(() => {
    if(authenticated) {
      router.push('/')
      return
    }
  }, [authenticated, router])

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault(); 
    try {
      console.log('submitHandler called', [email, password]);
      console.log('★★★★ End listing inner soldiers ★★★★')
      const res = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: {'Content-Type':'application/json'},
        body: JSON.stringify({email,password})
      })

      const data = await res.json()

      if(res.ok) {
        const user: UserInfo = {
          userName:data.username,
          id: data.id
        }

        localStorage.setItem('user_info', JSON.stringify(user))
        return router.push('/')
      }
    } catch (err) {
      console.log(err);
      
    }
  }

  return (
    
    <div className='flex items-center justify-center min-w-full min-h-screen'>
      <form className='flex flex-col md:w-1/5'>
        <div className='text-3xl font-bold text-center'>
          <span className='text-blue'>Welcome Back!</span>
        </div>

        <input
          placeholder='Email'
          className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
          value={email}
          onChange={(e)=> setEmail(e.target.value)}
        />
        <input
          placeholder='Password'
          type='password'
          className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
             value={password}
          onChange={(e)=> setPassword(e.target.value)}
        />
        <button className='p-3 mt-6 rounded-md bg-blue font-bold text-white'
        type='submit'
        onClick={submitHandler}>
          Login
        </button>
      </form>
    </div>
  );
};

export default Index;
