import React from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary, Avatar,
    Button,
    Container,
    Grid,
    Theme,
    Typography
} from "@material-ui/core";
import {logout} from "./AuthPage";
import {makeStyles} from "@material-ui/core/styles";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

const useStyles = makeStyles((theme: Theme) => ({
    indicator: {
        backgroundColor: '#fff',
    }
}));

const consents = [
    {
        name: "Financroo",
        id: "1234567",
        granted: "1 jan 2001",
        valid: "2 jan 2002",
        accounts: [
            {
                accountId: "233q432423",
                permissions: [
                    {
                        name: "permissionNamePermissionNamePermissionnanamam",
                        description: "descdescdescdescdescdesc descdesc descdescdescdescdesc descdesc descdescdesc descdesc descdesc"
                    },
                    {
                        name: "permission 1",
                        description: "desc"
                    },
                    {
                        name: "permission 1",
                        description: "desc"
                    }
                ]
            },
            {
                accountId: "233q43223",
                permissions: [
                    {
                        name: "permission 2",
                        description: "desc"
                    },
                    {
                        name: "permission 2",
                        description: "desc"
                    },
                    {
                        name: "permission 1",
                        description: "desc"
                    }
                ]
            }
        ]
    },
    {
        name: "Financroo2",
        id: "1234567",
        granted: "1 jan 2001",
        valid: "2 jan 2002",
        accounts: [
            {
                accountId: "233q432423",
                permissions: [
                    {
                        name: "permission 1",
                        description: "desc"
                    },
                    {
                        name: "permission 1",
                        description: "desc"
                    },
                    {
                        name: "permission 1",
                        description: "desc"
                    }
                ]
            },
            {
                accountId: "233q43223",
                permissions: [
                    {
                        name: "permission 2",
                        description: "desc"
                    },
                    {
                        name: "permission 2",
                        description: "desc"
                    },
                    {
                        name: "permission 1",
                        description: "desc"
                    }
                ]
            }
        ]
    }
]

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
    const isProgress = false;
    const classes = useStyles();

    return (
        <div style={{position: 'relative'}}>
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
                        <Tab label="Get started" value={'get-started'} style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="APIs" value={'apis'} style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Help Center" value={'help-center'} style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Connected Applications" value={'connected-applications'} style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Monitoring" value={'monitoring'} style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                    </Tabs>
                </Hidden>
                <div style={{flex: 1}}/>
                <Button
                    style={{color: '#fff'}}
                    onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
            </PageToolbar>
            {isProgress && <Progress/>}

            {!isProgress && (
                <>
                    <div style={{
                        background: "#F7F7F7",
                        height: 148,
                        marginTop: 64,
                        display: 'flex',
                        alignItems: 'center'
                    }}>
                        <Container>
                            <Typography variant={"h3"} color={"primary"}>Connected Applications</Typography>
                        </Container>
                    </div>
                    <Container style={{marginTop: 64}}>
                        <Grid container justify={"center"}>
                            <Grid item xs={8}>
                                {consents.map(consent => (
                                    <Accordion>
                                        <AccordionSummary
                                            expandIcon={<ExpandMoreIcon/>}
                                            aria-controls="panel1a-content"
                                            id="panel1a-header"
                                        >
                                            <div style={{
                                                display: "flex",
                                                justifyContent: "space-between",
                                                alignItems: "center",
                                                width: '100%'
                                            }}>
                                                <div style={{display: "flex", alignItems: "center"}}>
                                                    <Avatar>{consent.name[0]}</Avatar>
                                                    <Typography style={{marginLeft: 16}}>{consent.name}</Typography>
                                                </div>
                                                <div style={{flex: 1}}/>
                                                <Button variant={"contained"} color={"primary"} onClick={e => {
                                                    e.stopPropagation()
                                                }}>Revoke</Button>
                                            </div>
                                        </AccordionSummary>
                                        <AccordionDetails style={{flexDirection: "column", background: "#fcfcfc"}}>
                                            {consent.accounts.map(account => (
                                                <div style={{marginBottom: 32}}>
                                                    <Typography>
                                                        <strong>Account ID: {account.accountId}</strong>
                                                    </Typography>
                                                    <Typography style={{marginTop: 16, borderBottom: "1px solid #ccc", paddingBottom: 8, marginBottom: 16}}>Granted permissions:</Typography>
                                                    <Grid container spacing={2}>
                                                        {account.permissions.map(permission => (
                                                            <>
                                                                <Grid item xs={4}
                                                                      wrap={"wrap"}
                                                                      style={{textAlign: "right", wordWrap: "break-word"}}>
                                                                    <Typography><strong>{permission.name}</strong></Typography>
                                                                </Grid>
                                                                <Grid item xs={8}
                                                                      style={{textAlign: "left", wordWrap: "break-word"}}>
                                                                    <Typography>{permission.description}</Typography>
                                                                </Grid>
                                                            </>
                                                        ))}

                                                    </Grid>
                                                </div>
                                            ))}
                                        </AccordionDetails>
                                    </Accordion>
                                ))}
                            </Grid>
                        </Grid>

                    </Container>
                </>
            )}
        </div>
    )
};
