import React from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import backgroundLogin from "../assets/background-login.png";
import financrooLogo from "../assets/financroo-logo.svg";
import Grid from "@material-ui/core/Grid";
import {useFormFactory} from "../utils/forms/formFactory";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Button from "@material-ui/core/Button";

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    height: '100vh'
  },
  image: {
    width: '100%',
    height: '100%',
    backgroundImage: `url(${backgroundLogin})`,
    backgroundRepeat: 'no-repeat',
    backgroundPosition: 'center',
    backgroundSize: 'cover'
  },
  formContainerRoot: {
    [theme.breakpoints.down('sm')]: {
      padding: 16
    },
    [theme.breakpoints.up('sm')]: {
      padding: 48
    },
    [theme.breakpoints.up('lg')]: {
      padding: 82
    }
  }
}));

export default function Register ({onLogin}) {
  const classes = useStyles();
  const formFactory = useFormFactory({id: 'login-view'})

  return (
    <div className={classes.root}>
      <Grid container style={{height: '100%'}}>
        <Grid item sm={6} lg={7}>
          <div className={classes.image}/>
        </Grid>
        <Grid item xs={12} sm={6} lg={5}>
          <div className={classes.formContainerRoot}>
            <img alt="financroo logo" src={financrooLogo} style={{marginBottom: 24}}/>
            <Typography variant={'h4'} style={{marginBottom: 24}}>Create your Financroo account</Typography>
            {formFactory.createRequiredField({name: 'username', label: 'Username'})}
            {formFactory.createRequiredField({name: 'password', label: 'Password'})}

            {formFactory.createFormFooter({onSubmit: data => console.log(data), submitText: 'Create Account', submitButtonColor: 'secondary', buttonWidth: '100%'})}
            <Divider style={{margin: '24px 0'}}/>
            <Typography style={{textAlign: 'center'}}>By clicking Create account, you agree to our
              Terms of service and have read and acknowledge our Privacy Statement</Typography>
            <Typography style={{marginTop: 32, textAlign: 'center'}}>
              <strong>Already have an account? </strong>
              <Button onClick={onLogin}>Login</Button>
            </Typography>
          </div>
        </Grid>
      </Grid>
    </div>
  )
};
