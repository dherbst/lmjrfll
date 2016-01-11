+++
date = "2016-01-08T07:43:12-05:00"
draft = false
title = "register"

+++
<div ng-if="user.IsLoggedIn">
<h1>You have already registered!</h1>

You can go to your <a href="/expo/my_team/index.html">team page</a> to see your team details.
</div>

# Register for the Lower Merion Jr. First Lego League

<form method="post" action="">
  <p><label for="id_username">Username:</label> <input id="id_username" type="text" class="required" name="username" maxlength="30"></p>
  <p><label for="id_email">E-mail:</label> <input id="id_email" type="text" class="required" name="email" maxlength="75"></p>
  <p><label for="id_password1">Password:</label> <input id="id_password1" type="password" class="required" name="password1"></p>
  <p><label for="id_password2">Password (again):</label> <input id="id_password2" type="password" class="required" name="password2"></p>
  <p><input type="submit" value="Send activation email"></p>
</form>

You will receive an email with a link to confirm your email address.  If you do not get it in the next few minutes, please search your spam folder for "lowermerionjrfll.com".
