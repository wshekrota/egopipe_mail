# egopipe_mail
Read some emails to analyze in a data pipeline.

### A simple usecase for egopipe. What is a more universal example of raw data than email?

Assumption: you have a go environment already.

Clone this repository somewhere for your reference. 

https://github.com/wshekrota/egopipe_mail.git

You will need the mail.go and your_pipe_code.go
sources later in this setup.

Given the following package which reads mail from an imap server according to rfc thanks to Brian..

go get https://github.com/BrianLeishman/go-imap/

This will install that package for you. Then I took his simple example code from the README and 
modified it for the purpose. (this is the mail.go you just cloned) It will output something we can feed in logstash from filebeat.
In that way the logfile is managed and could be cron'ed for continuous use or just a one off.

**Install** filebeat perhaps on your laptop OS. Then go to directory /etc/filebeat and edit filebeat.yml
to indicate a logstash target of 172.17.0.3:5000. Installation will depend on your OS type.

In mail.go source...
**Configure** your account info..
![imap account configuration](/images/account.png)

Notice the lines that select what emails are of interest. (in this case they are inmail emails from linkedin)
![selection part of mail executable](/images/mailselect.png)

Once you have configured that part of the source compile mail.go.
'go build mail.go' 
Then **copy** the executable to your log directory /var/log so you can create the logfile.
'**mail > myfiles.log**'
The output of this will be a line of json for each email item.

Then **startup** filebeat..
'service filebeat start'

Now filebeat is in wait mode due to logstash not answering.

**Clone** the egopipe repo down and copy in the your_pipe_code.go from the egopipe_mail so you can make any changes.
https://github.com/wshekrota/egopipe.git
Least case copy the your_pipe_code.go from egopipe_mail cloned above to here.

'go build ./...' then copy the egopipe executable 'egopipe' to the Dockerfile directory below for logstash.

```
What is the your_pipe_code.go? It is the transform stage of the pipe. This constitutes the changes or amendments 
you make to the data. As coded the file included with mail changes the following for the log file named..

- decode json
- changes the @timestamp to the actual time the email was received
- compares 3 different strings to the text field of the email setting an indicator field. 
(so you may want to modify this to meet your email analysis need)
```

Now we **clone** the repo containing the Dockerfiles to start the Elasticstack components. 

https://github.com/wshekrota/egopipe_containers.git

cd to each directory ....apps/elastic (and do a 'docker build .') after assuring the files you need are there. elastic and kibana you can use with minor changes to the config files. Logstash you want to be sure the egopipe executable wass compilled and copied there. It will be placed in the container.

Here are some kibana images of linkedin data for my jobsearch.
![donut all data](/images/li.png)
Here are totals by day...
![daily counts](/images/dailymix.png)
This is the tranform code that enhanced the data. (your_pipe_code.go)
For your linkedin emails you may analyze the types of jobs differently? I was looking at 3 categories.
![transform code](/images/mailtran.png)
