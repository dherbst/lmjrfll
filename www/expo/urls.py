from django.conf.urls import patterns, include, url

urlpatterns = patterns(
    'expo.views',
    url(r'^$', 'expo_home', name='expo_home'),

)
