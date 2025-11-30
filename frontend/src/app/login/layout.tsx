'use client'
import LoadingPage from '@/components/common/loading/LoadingPage';
import StarsBackground from '@/components/common/particles/StartParticles';
import { useAuth } from '@/hooks/useAuth';
import { useRouter } from 'next/navigation';
import React, { ReactNode } from 'react';

export default function LoginLayout({ children }: { children: ReactNode }) {

    const { user} = useAuth();
    const router = useRouter();

    if (user) {

        router.replace('/manager/home');

        return <LoadingPage loadingText='Redirigiendo' />

    } else {

        return <section className="relative">
            <div className="absolute">
                <StarsBackground />
            </div>
            <div className="relative min-h-screen flex items-center justify-center">
                {children}
            </div>
        </section>

    }

}
