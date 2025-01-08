import type { Metadata } from 'next';
import { AppRouterCacheProvider } from '@mui/material-nextjs/v15-appRouter';
import CssBaseline from '@mui/material/CssBaseline';
import { Roboto } from 'next/font/google';

import ReduxProvider from '@/store/ReduxProvider';
import {
  chartsCustomizations,
  dataGridCustomizations,
  datePickersCustomizations,
  treeViewCustomizations
} from '@/theme/customizations';

import './globals.css';
import AppTheme from '@repo/theme/AppTheme';

const roboto = Roboto({
  weight: ['300', '400', '500', '700'],
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-roboto'
});

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app'
};

const xThemeComponents = {
  ...chartsCustomizations,
  ...dataGridCustomizations,
  ...datePickersCustomizations,
  ...treeViewCustomizations
};

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={roboto.variable}>
        <ReduxProvider>
          <AppRouterCacheProvider>
            <AppTheme themeComponents={xThemeComponents}>
              <CssBaseline enableColorScheme />
              {children}
            </AppTheme>
          </AppRouterCacheProvider>
        </ReduxProvider>
      </body>
    </html>
  );
}
