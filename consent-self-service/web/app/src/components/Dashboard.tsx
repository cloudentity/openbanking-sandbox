import React, {useEffect, useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {Button, Container, Grid, Theme, Typography} from "@material-ui/core";
import {logout} from "./AuthPage";
import {makeStyles} from "@material-ui/core/styles";
import {api} from "../api/api";
import noAccountEmptyState from "./no-accounts-empty-state.svg";
import ConfirmationDialog from "./ConfirmationDialog";
import ApplicationCard from "./ApplicationCard";

const useStyles = makeStyles((theme: Theme) => ({
    indicator: {
        backgroundColor: '#fff',
    },
    expandIcon: {
        position: 'absolute', right: 32, top: 24, color: "#006580"
    }
}));

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
    const [isProgress, setProgress] = useState(true);
    const [clientConsents, setClientConsents] = useState<any>([]);
    const [revokeDialog, setRevokeDialog] = useState<any>(null);
    const classes = useStyles();

    useEffect(() => {
        setProgress(true);
        api.getConsents()
            .then(res => setClientConsents(res.client_consents))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }, []);

    const handleRevoke = id => {
        setProgress(true);
        api.deleteConsent({id})
            .then(api.getConsents)
            .then(res => setClientConsents(res.client_consents))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }

    return (
        <div>
            <PageToolbar>
                <Hidden mdUp>
                    <IconButton edge="start" color="inherit" aria-label="menu">
                        <MenuIcon/>
                    </IconButton>
                </Hidden>
                <Hidden smDown>
                    <Tabs
                        value={'connected-applications'}
                        onChange={() => {
                        }}
                        classes={{
                            indicator: classes.indicator
                        }}
                        aria-label="menu tabs"
                        style={{height: 64}}
                    >
                        <Tab label="Connected Applications" value={'connected-applications'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                    </Tabs>
                </Hidden>
                <div style={{flex: 1}}/>
                <Button
                    style={{color: '#fff'}}
                    onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
            </PageToolbar>
            <div style={{marginTop: 64, position: "relative"}}>
                {isProgress && <Progress/>}

                {!isProgress && (
                    <>
                        {clientConsents.length > 0 && (
                            <div style={{
                                background: "#F7F7F7",
                                height: 148,
                                display: 'flex',
                                alignItems: 'center'
                            }}>
                                <Container>
                                    <Typography variant={"h3"} color={"primary"}>Connected Applications</Typography>
                                </Container>
                            </div>
                        )}
                        <Container style={{marginTop: 64}}>
                            <Grid container justify={"center"}>
                                <Grid item xs={8}>
                                    {clientConsents.length === 0 && (
                                        <div style={{textAlign: "center", marginTop: 64}}>
                                            <Typography variant={"h3"} style={{color: "#626576"}}>No connected
                                                accounts</Typography>
                                            <Typography style={{marginTop: 12, color: "#A0A3B5"}}>You haven’t connected
                                                any accounts yet
                                                to manage
                                                access</Typography>
                                            <img src={noAccountEmptyState} style={{marginTop: 64}} alt={"empty state"}/>
                                        </div>
                                    )}
                                    {clientConsents.map(clientConsent => (
                                        <ApplicationCard client={clientConsent} onRevokeConsent={consent => setRevokeDialog(consent)} key={clientConsent.id}/>
                                    ))}
                                </Grid>
                            </Grid>

                        </Container>
                    </>
                )}
            </div>
            {revokeDialog && (
                <ConfirmationDialog
                    title="Revoke application access"
                    content="Are you sure you want to revoke access to your accounts for this application?"
                    confirmText="Yes, Revoke Access"
                    warningItems={[
                        'Your connected application will not be able to access your bank accounts.',
                        'You will have to authorize the application again if you would like to use it with your bank in the future.'
                    ]}
                    onCancel={() => setRevokeDialog(null)}
                    onConfirm={() => {
                        handleRevoke(revokeDialog.consent_id);
                        setRevokeDialog(null);
                    }}
                />
            )}
        </div>
    )
};
