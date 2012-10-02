from django.http import HttpResponse
from django.shortcuts import render_to_response, get_object_or_404
from django.template import loader, Context, RequestContext

def home(request, template='lml_index.html'):
    
    return render_to_response(
        template,
        {},
        context_instance=RequestContext(request))


    
