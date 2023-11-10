import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import StoreProvider from './components/StoreProvider';

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'FinTech',
  description: 'Financial App',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <StoreProvider>
      <body className={inter.className}>
        {children}
        <ToastContainer/>
        </body>
        </StoreProvider>
    </html>
  )
}
