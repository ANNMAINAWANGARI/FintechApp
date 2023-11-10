'use client'
import Image from 'next/image'
import MainLayout from './components/MainLayout'

export default function Home() {
  return (
    
    <main className="flex min-h-screen flex-col px-16 py-5 font-bold justify-between bg-black text-white">
      <div className='text-2xl subpixel-antialiased tracking-wide'>FinTech </div>
      <section className='flex flex-col items-center justify-items-center mt-20 gap-3'>
        <h1 className='text-6xl font-extrabold'>Master your <span className='text-fuchsia-800'>Finance</span></h1>
        <p>tracks your accounts,manage expenses,and budgets like a pro!</p>
        <button className='bg-fuchsia-800 rounded-lg p-2 my-5'>Signup for Waitlist</button>
        <Image
              src="/cover.webp"
              alt="Fintech Cover"
              className="mt-5"
              width={500}
              height={100}
              priority
            />
      </section>
    </main>
    
  )
}
