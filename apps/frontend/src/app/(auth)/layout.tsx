import { Container } from '@mui/material';

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <Container component="main" maxWidth="xs">
      {children}
    </Container>
  );
}
