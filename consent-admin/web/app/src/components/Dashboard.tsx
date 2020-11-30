import React, {useEffect, useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {Button, Theme} from "@material-ui/core";
import {logout} from "./AuthPage";
import {makeStyles} from "@material-ui/core/styles";
import {api} from "../api/api";
import ConfirmationDialog from "./ConfirmationDialog";
import {Route} from "react-router-dom";
import AccountsContainer from "./AccountsContainer";
import AccountsDetailsContainer from "./AccountsDetailsContainer";

const useStyles = makeStyles((theme: Theme) => ({
    indicator: {
        backgroundColor: '#fff',
    }
}));

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
    const [isProgress, setProgress] = useState(true);
    const [data, setData] = useState<any>([]);
    const [revokeDialog, setRevokeDialog] = useState<any>(null);
    const classes = useStyles();

    useEffect(() => {
        setProgress(true);
        api.getConsents()
            .then(res => setData(res.consents))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }, []);

    const handleRevoke = id => {
        setProgress(true);
        api.deleteConsent({id})
            .then(api.getConsents)
            .then(res => setData(res.consents))
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
                        <Tab label="Get started" value={'get-started'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="APIs" value={'apis'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Help Center" value={'help-center'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Connected Applications" value={'connected-applications'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Monitoring" value={'monitoring'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                    </Tabs>
                </Hidden>
                <div style={{flex: 1}}/>
                <Button
                    style={{color: '#fff'}}
                    onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
            </PageToolbar>
            <div style={{marginTop: 128, position: "relative"}}>
                {isProgress && <Progress/>}

                {!isProgress && (
                    <div>
                        <Route exact path="/" render={() => <AccountsContainer/>}/>
                        <Route exact path="/:id" render={() => <AccountsDetailsContainer/>}/>
                    </div>
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
