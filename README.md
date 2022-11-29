# What we want?

Use "right shift" to switch the input source instead of "ctrl + space".

That would be much more efficient for typing.

# What we need to do?

1. run **"go install"**

2. move **${GOROOT}/switchInputSource** to **~/Library/LaunchAgents/**

3. copy **org.terry.sis.plist** to **~/Library/LaunchAgents/** (don't forget to change the path for switchInputSource in the plist file)

4. relogin