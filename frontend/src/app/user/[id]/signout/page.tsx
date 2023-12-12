"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import axios from "@/api/axiosConfig";

import { useAuth } from "@/app/userContext";
import { Button, Typography, Box, Paper } from "@mui/material";

export default function SignOutPage() {
  const { logout, user } = useAuth();
  const router = useRouter();
  const onSignOut = async () => {
    const params = {
      id: user?.id,
    };
    window.localStorage.removeItem("session_token");
    logout();
    await axios.post("/logout", params);
    router.push("/");
  };
  return (
    <Box sx={{ display: "flex", alignItems: "center", height: "100vh" }}>
      <Paper
        elevation={3}
        sx={{ p: 4, borderRadius: 2, maxWidth: 400, mx: "auto" }}
      >
        <Box sx={{ textAlign: "center" }}>
          <Typography variant="h5" component="h1" sx={{ mb: 2 }}>
            Are You Sure You Want To Sign Out?
          </Typography>

          <Button
            variant="contained"
            color="primary"
            size="large"
            sx={{ px: 4 }}
            onClick={onSignOut}
          >
            Go SignOut
          </Button>
        </Box>
      </Paper>
    </Box>
  );
}
