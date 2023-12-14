"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "@/api/axiosConfig";

import { useAuth } from "@/app/userContext";
import { Button, Typography, Box, Paper } from "@mui/material";

import SnackBar from "@/components/SnackBar";

export default function SignOutPage() {
  const [open, setOpen] = useState(false);
  const { logout, user } = useAuth();
  const router = useRouter();
  const onSignOut = async () => {
    if (!user) {
      return router.push("/");
    }
    const params = {
      id: user?.id,
    };
    try {
      const res = await axios.post("/logout", params);
      if (res.status !== 200) {
        return console.log("failed");
      }
      logout();
      router.push("/");
    } catch (error) {
      setOpen(true);
      console.log(error);
    }
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
      <SnackBar
        open={open}
        setOpen={setOpen}
        severity="error"
        text="failed to logout"
      ></SnackBar>
    </Box>
  );
}
