Output kuzco_demo.gif

Set WindowBar Colorful
Set Margin 20
Set Framerate 60
Set MarginFill "#000000"
Set TypingSpeed 75ms
Set BorderRadius 10
Set Width 1300
Set Height 650
Set FontSize 12

Type@50ms "# Enhance your Terraform configurations with intelligent analysis powered by local LLMs"
Sleep 1s
Enter
Type "# Lets try it out!"
Sleep 1s
Enter
Hide
Type "alias kuzco=./kuzco
Enter
Type "clear"
Enter
Show
Type@80ms "kuzco --file ../test/s3-bucket.tf --model PE3B:latest --prompt 'Provide security specific recommendations for the creation of this S3 bucket'"
Enter
Sleep 30s
