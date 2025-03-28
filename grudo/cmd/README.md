# Groups.io large org tools

Groups.io is a mailing list service. The service does a great job of managing mailing lists but the permissions model 
for groups is centered around a single mailing list. This means that in order to have the elevated permissions
(be an 'Owner' in groups.io terminology) you MUST be a member of the mailing list. For end-users in an organization 
that are managing hundreds of mailing lists, the absence of a admin role is awkward but thankfully groups.io does have a 
stable REST API which allows the opportunity to fake creating administrative users using the basic permissions model of 
Groups.io

The code in this repo offers a simple workaround to carve out notional administrative roles by having one user take on 
the subscriptions of another user.

I wrote this code for the CNCF but if you have or can acquire go programming skills you should be able to make use of 
the code you find here.    

## Usage
You will need install go to run this code.

Everything you can do using gio-admin will operate from the permissions and privilege you have from your existing list
subscriptions. TODO Verify that this statement is correct.



### List your Subscriptions

### Transfer Subscriptions

    --srcPass=***** \
    --cmd=xferSubs \
    --destEmail=rkleinhans@contractor.linuxfoundation.org \ 
    --baseUrl=https://lists.cncf.io \