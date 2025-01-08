'use client';

import { alpha, Box, Stack } from '@mui/material';

import AppNavbar from '@/components/AppNavbar';
import Header from '@/components/Header';
import SideMenu from '@/components/SideMenu';

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <Box sx={{ display: 'flex' }}>
        <SideMenu />
        <AppNavbar />
        {/* Main content */}
        <Box
          component="main"
          sx={(theme) => ({
            flexGrow: 1,
            backgroundColor: alpha(theme.palette.background.default, 1),
            overflow: 'auto'
          })}>
          <Stack
            spacing={2}
            sx={{
              alignItems: 'center',
              mx: 3,
              pb: 5,
              mt: { xs: 8, md: 0 }
            }}>
            <Header />
            {children}
          </Stack>
        </Box>
      </Box>
    </>
  );
}
