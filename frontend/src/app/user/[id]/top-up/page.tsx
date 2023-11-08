"use client";
import React, { useState } from "react";
import { Button, TextField, Typography } from "@mui/material";
import Image from "next/image";
import CustomSnackBar, { Severity } from "@/components/SnackBar";
import EWalletPic from "../../../../../public/e-wallet.png";
import { SnackBarTextErr, SnackBarTextSuccess } from "./text";
import { useAuth } from "@/app/userContext";

const TopUp = () => {
  const [amount, setAmount] = useState("");
  const [errOpen, setErrOpen] = useState<boolean>(false);
  const [successOpen, setSuccessOpen] = useState<boolean>(false);
  const { user } = useAuth();
  console.log(user);

  const handleTopUp = () => {
    // Here you can add the logic to top up the user's account

    if (!amount) return;
    if (!Number(amount)) {
      setErrOpen(true);
    }
    console.log(`Top up with amount: ${amount}`);
    setSuccessOpen(true);
  };

  return (
    <main
      style={{
        margin: "30px",
        textAlign: "center",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography>This is Top up page</Typography>
      <Image src={EWalletPic} alt="e-wallets" loading="lazy" width={300} />
      <Typography>Valance: 30</Typography>
      <TextField
        label="Amount"
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
        error={Number(amount) ? false : !amount ? false : true}
      />
      <Button onClick={handleTopUp}>Top Up</Button>
      <CustomSnackBar
        open={errOpen}
        setOpen={setErrOpen}
        severity={Severity.ERROR}
        text={SnackBarTextErr}
      />
      <CustomSnackBar
        open={successOpen}
        setOpen={setSuccessOpen}
        severity={Severity.SUCCESS}
        text={SnackBarTextSuccess}
      />
    </main>
  );
};

export default TopUp;
