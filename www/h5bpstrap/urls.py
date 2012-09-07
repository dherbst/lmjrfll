from django.conf.urls import patterns, include, url

urlpatterns = patterns(
    '',
    url(r'^$', 'h5bpstrap.views.home', name='home'),
    )
